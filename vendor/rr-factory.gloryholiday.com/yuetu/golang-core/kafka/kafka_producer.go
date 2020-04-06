package kafka

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/Shopify/sarama"
	"rr-factory.gloryholiday.com/yuetu/golang-core/logger"
)

const (
	running int32 = 0
	stopped int32 = 1
)

type Topic string
type PK string

type KafkaMessenger interface {
	Topic() Topic
	PK() PK
	Payload() ([]byte, error)
}

type KafkaProducer interface {
	Send(KafkaMessenger) error
}

type kafkaProducer struct {
	alive    int32
	timeout  time.Duration
	ctx      context.Context
	producer sarama.AsyncProducer
}

var KafkaNotAlive error = errors.New("kafka is not alive")
var kafkaEnabled = true

func DisableKafka() {
	kafkaEnabled = false
}

func EnableKafka() {
	kafkaEnabled = true
}

func (kp *kafkaProducer) Send(messenger KafkaMessenger) error {
	if kp.isAlive() {
		var err error
		pk := string(messenger.PK())
		topic := string(messenger.Topic())

		key := sarama.StringEncoder(pk)
		payload, err := messenger.Payload()
		if err != nil {
			return err
		}
		value := sarama.ByteEncoder(payload)
		select {
		case kp.producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: key, Value: value}:
			logger.Debug(logger.Message("Message sent[%s -> %s]", topic, key))
		case <-time.After(kp.timeout):
			err = errors.New("Timeout [500ms] to send kafka message")
			logger.ErrorNt("message pk: "+pk, err)
		}

		return err
	}

	if kp.isEnabled() {
		logger.InfoNt("Kafka is disabled")
	}

	if kp.isRunning() {
		logger.WarnNt("Kafka producer has been stopped")
	}

	return KafkaNotAlive
}

func (kp *kafkaProducer) stats() {
	go func() {
		defer func() {
			kp.stop()
			if err := kp.producer.Close(); err != nil {
				logger.WarnNt(logger.Message("Failed to close async kafka producer: %s", err.Error()))
			}
		}()

		for {
			select {
			case err := <-kp.producer.Errors():
				logger.WarnNt(logger.Message("Failed to produce message: %s", err.Error()))
			case <-kp.ctx.Done():
				logger.InfoNt("Stop KafkaProducer by context notification")
				return
			}
		}
	}()
}

func (kp *kafkaProducer) isAlive() bool {
	return kp.isEnabled() && kp.isRunning()
}

func (kp *kafkaProducer) isEnabled() bool {
	return kafkaEnabled
}

func (kp *kafkaProducer) isRunning() bool {
	return atomic.LoadInt32(&kp.alive) == running
}

func (kp *kafkaProducer) stop() {
	atomic.StoreInt32(&kp.alive, stopped)
}

func newKafkaClient(brokers []string) (sarama.Client, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_1_0_0
	saramaConfig.Producer.Compression = sarama.CompressionLZ4

	return sarama.NewClient(brokers, saramaConfig)
}

func newKafkaProcuder(ctx context.Context, brokers []string) (KafkaProducer, error) {
	client, err := newKafkaClient(brokers)
	if err != nil {
		return nil, err
	}

	return newKafkaProducerByClient(ctx, client)
}

func newKafkaProducerByClient(ctx context.Context, client sarama.Client) (KafkaProducer, error) {
	producer, err := sarama.NewAsyncProducerFromClient(client)

	if err != nil {
		return nil, err
	}

	kp := &kafkaProducer{
		producer: producer,
		timeout:  500 * time.Millisecond,
		ctx:      ctx,
	}

	kp.stats()

	return kp, nil
}

type producerConfig struct {
	brokers             []string
	numberOfProducers   int
	dispatchingStrategy func() int
}

type producerOption func(*producerConfig)

type multiplexKafkaProducer struct {
	config    *producerConfig
	producers []KafkaProducer
}

func newRRDispathingStrategy(totalProducerNumber int) func() int {
	return func() int {
		if totalProducerNumber == 1 {
			return 0 // there's no need to choose the target producer based on random number if there's only one available
		}
		return rand.Intn(totalProducerNumber)
	}
}

func newProducerConfig() *producerConfig {
	return &producerConfig{
		numberOfProducers: 2,
	}
}

func WithProducerNumber(number int) producerOption {
	return func(cfg *producerConfig) {
		cfg.numberOfProducers = number
	}
}

func NewMultiplexKafkaProducer(ctx context.Context, brokers []string, options ...producerOption) (KafkaProducer, error) {
	cfg := newProducerConfig()
	cfg.brokers = brokers

	for _, option := range options {
		option(cfg)
	}

	return NewMultiplexKafkaProducerFromConfig(ctx, cfg)
}

func NewMultiplexKafkaProducerFromConfig(ctx context.Context, cfg *producerConfig) (KafkaProducer, error) {
	client, err := newKafkaClient(cfg.brokers)
	if err != nil {
		return nil, err
	}

	producers := []KafkaProducer{}

	for i := 0; i < cfg.numberOfProducers; i++ {
		producer, err := newKafkaProducerByClient(ctx, client)
		if err != nil {
			logger.ErrorNt("Failed to new async producer", err)
			continue
		}
		producers = append(producers, producer)
	}

	logger.InfoNt(logger.Message("%d producers created, wanted %d", len(producers), cfg.numberOfProducers))

	if len(producers) == 0 {
		return nil, errors.New("No producer created successfully")
	}

	cfg.dispatchingStrategy = newRRDispathingStrategy(len(producers))
	return &multiplexKafkaProducer{
		producers: producers,
		config:    cfg,
	}, nil
}

func (mkp *multiplexKafkaProducer) Send(messenger KafkaMessenger) error {
	if !kafkaEnabled {
		return nil
	}
	producerIndex := mkp.config.dispatchingStrategy()
	logger.Debug(logger.Message("Producer %d selected", producerIndex))

	if producerIndex >= len(mkp.producers) {
		return fmt.Errorf("dispatching strategy seems not right, target producer index out of bound. expected: %d, actual: %d", producerIndex, len(mkp.producers))
	}

	return mkp.producers[producerIndex].Send(messenger)
}
