/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 10:56
 */
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"gochat/api"
	"gochat/connect"
	"gochat/logic"
	"gochat/pkg/exec"
	"gochat/site"
	"gochat/task"
)

func main() {
	var module string
	var other string
	flag.StringVar(&module, "module", "", "assign run module")
	flag.StringVar(&other, "other", "", "other parameter")
	flag.Parse()
	fmt.Printf("start run %s module\n", module)
	switch module {
	case "logic":
		logic.New().Run()
	case "connect_websocket":
		connect.New().Run()
	case "connect_tcp":
		connect.New().RunTcp()
	case "task":
		task.New().Run()
	case "api":
		api.New().Run()
	case "site":
		site.New().Run()
	case "all":
		withServiceCmd, _ := strconv.ParseBool(other)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		exec.StartAll(ctx, withServiceCmd)
	default:
		fmt.Println("exiting, module param error!")
		return
	}
	fmt.Printf("run %s module done!\n", module)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	fmt.Println("Server exiting")
}
