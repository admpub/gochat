package exec

import (
	"os"
	"time"

	"github.com/webx-top/com"
)

func StartAll(withServiceCmd bool) error {
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
	if withServiceCmd {
		for _, cmd := range serviceCmds {
			com.RunCmdStrWithWriter(cmd)
		}
		time.Sleep(5 * time.Second)
	}
	for _, cmd := range moduleCmds {
		com.RunCmdStrWithWriter(cmd)
	}
	return nil
}
