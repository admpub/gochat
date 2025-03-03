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
	var other bool
	flag.StringVar(&module, "module", "", "assign run module")
	flag.BoolVar(&other, "other", false, "other parameter")
	flag.Parse()
	fmt.Printf("start run %s module\n", module)

	/*
	          [connect] (websocket or tcp)
	         /                                \
	   [site]                                  [task]<RPC> (send message)
	         \                                /
	   	      [api]<RPC> —— [logic]<Queue> ——

	*/
	switch module {

	// -------------------- connect begin
	// 含房间管理逻辑，比如：出入房间
	// 将 authToken 通过 RPC 传递给 logic 查询 userId。通过 userId 来核实用户登录状态

	case "connect_websocket": // Websocket+RPC Server
		go connect.New().Run()
	case "connect_tcp": // TCP+RPC Server
		connect.New().RunTcp()

	// -------------------- connect end

	case "task": // (Redis)Queue Consumer + (connect)RPC Caller
		task.New().Run()
	case "logic": // RPC Server + (Redis)Queue Producer
		// RPC服务包含对数据库的操作：用户账号的注册、登录等
		logic.New().Run()
	case "api": // API Server + (logic)RPC Caller
		api.New().Run()
	case "site": // Frontend static file server ( Client: API Caller + (websocket/tcp)Connect Caller )
		go site.New().Run()

	case "all":
		withServiceCmd := other
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
	fmt.Println("[" + module + "]Server exiting")
}
