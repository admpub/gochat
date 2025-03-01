package exec

import (
	"context"
	"os"
	"time"

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
	if withServiceCmd {
		for _, cmd := range serviceCmds {
			com.RunCmdStrWriterWithContext(ctx, cmd)
		}
		time.Sleep(5 * time.Second)
	}
	for _, cmd := range moduleCmds {
		com.RunCmdStrWriterWithContext(ctx, cmd)
	}
	return nil
}
