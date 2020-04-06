package parallel

import (
	"fmt"
	"reflect"
	"sync"
)

type ParallelTaskStatus int

const (
	Pending ParallelTaskStatus = iota
	Processing
	Canceled
	Completed
	Failed
)

type ParallelTask struct {
	Index  int         // For Slice iterators
	Key    interface{} // For map iterators
	Val    interface{}
	Err    error
	Status ParallelTaskStatus
}

type ParallelTaskProducer interface {
	Produce() (*ParallelTask, bool)
}

type ParallelTaskConsumer func(*ParallelTask) error

type ParallelExecutor interface {
	Execute(ParallelTaskProducer, ParallelTaskConsumer)
}

type executorConfig struct {
	maxConcurrency int32
}

type syncedParallelExecutor struct {
	sync.WaitGroup
	config          *executorConfig
	concurrencyChan chan struct{}
}

type sliceTaskProducer struct {
	slice    reflect.Value
	size     int
	curIndex int
}

type channelTaskProducer struct {
	ch       reflect.Value
	curIndex int
}

type consumer struct {
	handler func(interface{}) error
}

type executorOption func(*executorConfig)

func WithMaxConcurrency(max int32) executorOption {
	return func(config *executorConfig) {
		config.maxConcurrency = max
	}
}

func (task *ParallelTask) String() string {
	return fmt.Sprintf("Index: %d, Key: %v, Val: %v, Err: %v, Status: %v", task.Index, task.Key, task.Val, task.Err, task.Status)
}

func SliceTaskProducer(slice interface{}) ParallelTaskProducer {
	s := reflect.ValueOf(slice)

	return &sliceTaskProducer{
		slice: s,
		size:  s.Len(),
	}
}

func ChannelTaskProducer(ch interface{}) ParallelTaskProducer {
	s := reflect.ValueOf(ch)

	return &channelTaskProducer{
		ch: s,
	}
}

type mapTaskProducer struct {
	mapIter *reflect.MapIter
}

func MapTaskProducer(m interface{}) ParallelTaskProducer {
	s := reflect.ValueOf(m)

	return &mapTaskProducer{s.MapRange()}
}

func (producer *mapTaskProducer) Produce() (*ParallelTask, bool) {
	if producer.mapIter.Next() {
		return &ParallelTask{
			Key:    producer.mapIter.Key().Interface(),
			Val:    producer.mapIter.Value().Interface(),
			Status: Pending,
		}, true
	}

	return nil, false
}

func (producer *sliceTaskProducer) Produce() (*ParallelTask, bool) {
	if producer.curIndex >= producer.size {
		return nil, false
	}

	curIndex := producer.curIndex
	producer.curIndex += 1

	return &ParallelTask{
		Index:  curIndex,
		Val:    producer.slice.Index(curIndex).Interface(),
		Status: Pending,
	}, true
}

func (producer *channelTaskProducer) Produce() (*ParallelTask, bool) {
	x, ok := producer.ch.Recv()
	defer func() {
		producer.curIndex += 1
	}()

	if ok {
		return &ParallelTask{
			Index:  producer.curIndex,
			Val:    x.Interface(),
			Status: Pending,
		}, true
	}

	return nil, false
}

func NewSyncedParallelExecutor(options ...executorOption) ParallelExecutor {
	config := &executorConfig{}

	for _, option := range options {
		option(config)
	}

	executor := &syncedParallelExecutor{
		config: config,
	}

	if executor.isConcurrencyLimited() {
		executor.concurrencyChan = make(chan struct{}, config.maxConcurrency)
	}

	return executor
}

func (executor *syncedParallelExecutor) isConcurrencyLimited() bool {
	return executor.config.maxConcurrency > 0
}

func (executor *syncedParallelExecutor) startTask() {
	executor.Add(1)
	if executor.isConcurrencyLimited() {
		executor.concurrencyChan <- struct{}{}
	}
}

func (executor *syncedParallelExecutor) finishTask() {
	executor.Done()
	if executor.isConcurrencyLimited() {
		<-executor.concurrencyChan
	}
}

func (executor *syncedParallelExecutor) Execute(producer ParallelTaskProducer, consumer ParallelTaskConsumer) {
	for {
		task, ok := producer.Produce()
		if ok {
			executor.startTask()
			go func(t *ParallelTask) {
				var err error
				defer func() {
					if r := recover(); r != nil {
						err = fmt.Errorf("Recover: %v", r)
					}
					if err != nil {
						task.Status = Failed
						task.Err = err
					} else {
						task.Status = Completed
					}
					executor.finishTask()
				}()
				task.Status = Processing

				err = consumer(t)
			}(task)
		} else {
			break
		}
	}
	executor.Wait()
}
