package concurrency

import (
	"fmt"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
	"rr-factory.gloryholiday.com/yuetu/golang-core/config"
	"rr-factory.gloryholiday.com/yuetu/golang-core/logger"
)

const (
	etcdConfigDir string = "etcd.config.dir"
)

func getPemFile(pemFileName string) string {
	configDir := config.GetString(etcdConfigDir)
	if len(configDir) == 0 {
		return fmt.Sprintf("/etc/etcd/%s/%s", config.GetProfile(), pemFileName)
	}

	return fmt.Sprintf("%s/%s", strings.TrimRight(configDir, "/"), pemFileName)
}

func getEtcdTlsCertFile() string {
	return getPemFile("client.pem")
}

func getEtcdTlsKeyFile() string {
	return getPemFile("client-key.pem")
}

func getEtcdTlsCaFile() string {
	return getPemFile("ca.pem")
}

func isTlsEnabled() bool {
	return !config.IsLocal()
}

func NewEtcdClient(endpoints []string) *clientv3.Client {
	etcdConfig := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	}
	if isTlsEnabled() {
		logger.InfoNt("Connecting etcd with TLS enabled.")
		tlsInfo := transport.TLSInfo{
			CertFile:      getEtcdTlsCertFile(),
			KeyFile:       getEtcdTlsKeyFile(),
			TrustedCAFile: getEtcdTlsCaFile(),
		}
		tlsConfig, err := tlsInfo.ClientConfig()

		if err != nil {
			logger.Fatal(logger.Message("Failed to create tls config. Error: %s", err.Error()))
		}

		etcdConfig.TLS = tlsConfig
	}
	etcd, err := clientv3.New(etcdConfig)
	if err != nil {
		logger.Fatal(logger.Message("Failed to connect etcd endpoints: %s. Error: %s", strings.Join(endpoints, ","), err.Error()))
	}

	logger.InfoNt(logger.Message("Connected etcd endpoints: %s", strings.Join(endpoints, ",")))

	return etcd
}
