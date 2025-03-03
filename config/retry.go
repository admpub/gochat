package config

import (
	"time"

	"github.com/sirupsen/logrus"
)

func Retry(name string, f func() (bool, error), err error, max int) error {
	var retry bool
	for i := 0; i < 10; i++ {
		logrus.Warnf("%s, err: %v. please wait to retry", name, err)
		time.Sleep(time.Second * 2)
		logrus.Warnf("[retry:%d]%s", i+1, name)
		retry, err = f()
		if err == nil || !retry {
			return err
		}
	}
	return err
}
