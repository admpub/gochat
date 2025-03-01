/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 18:22
 */
package task

import (
	"errors"
	"runtime"

	"gochat/config"

	"github.com/sirupsen/logrus"
)

type Task struct {
}

func New() *Task {
	return new(Task)
}

func (task *Task) Run() {
	//read config
	taskConfig := config.Conf.Task
	runtime.GOMAXPROCS(taskConfig.TaskBase.CpuNum)
	//read from redis queue
	if err := task.InitQueueRedisClient(); err != nil {
		logrus.Panicf("task init publishRedisClient fail, err: %s", err.Error())
	}
	//rpc call connect layer send msg
	if err := task.InitConnectRpcClient(); err != nil {
		if errors.Is(err, ErrNotFoundETCDServer) {
			err = config.Retry(`task init InitConnectRpcClient`, func() (bool, error) {
				err := task.InitConnectRpcClient()
				if err != nil {
					return errors.Is(err, ErrNotFoundETCDServer), err
				}
				return false, err
			}, err, 10)
			if err == nil {
				goto END
			}
		}
		logrus.Panicf("task init InitConnectRpcClient fail, err: %s", err.Error())
	}

END:
	//GoPush
	task.GoPush()
}
