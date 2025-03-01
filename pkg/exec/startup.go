package exec

import (
	"context"
	"os"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/smallnest/rpcx/log"
	"github.com/webx-top/com"
)

func StartAll(ctx context.Context, withServiceCmd bool) error {
	bin := os.Args[0]
	serviceCmds := []string{
		`etcd`,
		`redis-server`,
	}
	moduleCmds := []string{
		bin + ` -module logic`,
		bin + ` -module connect_websocket`,
		bin + ` -module task`,
		bin + ` -module api`,
		bin + ` -module site`,
	}
	cmds := make([]*exec.Cmd, 0, len(serviceCmds)+len(moduleCmds))
	if withServiceCmd {
		for _, cmd := range serviceCmds {
			logrus.Infof(`startup service: %s`, cmd)
			c := com.RunCmdStrWriterWithContext(ctx, cmd)
			cmds = append(cmds, c)
		}
		time.Sleep(5 * time.Second)
	}
	for _, cmd := range moduleCmds {
		logrus.Infof(`startup cmd: %s`, cmd)
		c := com.RunCmdStrWriterWithContext(ctx, cmd)
		cmds = append(cmds, c)
	}
	<-ctx.Done()
	println(`~~~~~~~~~~~~~~~~~~~~~~~~~~~~~shutdown~~~~~~~~~~~~~~~~~~~~~~~~`)
	for _, c := range cmds {
		if c.Process != nil {
			c.Process.Kill()
			log.Warnf(`kill %s`, c.Args[0])
		}
	}
	return nil
}
