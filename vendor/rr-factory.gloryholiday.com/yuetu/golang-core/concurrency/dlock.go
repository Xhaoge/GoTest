package concurrency

import (
	"context"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"rr-factory.gloryholiday.com/yuetu/golang-core/logger"
)

type DistributedTask struct {
	Ctx      context.Context
	TraceId  string
	Name     string
	LockPath string
	Task     func() error
}

type DistributedExecutor interface {
	Execute(*DistributedTask) error
}

type distributedExecutor struct {
	etcd *clientv3.Client
}

func NewEtcdDistributedExecutor(etcdEndpoints []string) *distributedExecutor {
	return &distributedExecutor{
		etcd: NewEtcdClient(etcdEndpoints),
	}
}

func (executor *distributedExecutor) Execute(task *DistributedTask) (err error) {
	logger.Info(task.TraceId, logger.Message("Start to execute task[%s]", task.Name))
	ctx, timeout := context.WithTimeout(task.Ctx, 5*time.Second)
	defer timeout()
	se, err := concurrency.NewSession(executor.etcd, concurrency.WithContext(ctx))
	if err != nil {
		return err
	}
	defer se.Close()

	mutex := concurrency.NewMutex(se, task.LockPath)
	if err := mutex.Lock(task.Ctx); err != nil {
		return err
	}
	logger.Info(task.TraceId, logger.Message("Acquired distributed lock [%s] for task [%s]", task.LockPath, task.Name))
	defer func() {
		err = mutex.Unlock(task.Ctx)
		logger.Info(task.TraceId, logger.Message("Released distributed lock [%s] and task [%s] finished", task.LockPath, task.Name))
	}()

	return task.Task()
}
