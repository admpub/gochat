/**
 * Created by lock
 * Date: 2019-10-06
 * Time: 22:46
 */
package rpc

import (
	"context"
	"gochat/config"
	"gochat/proto"
	"sync"
	"time"

	"github.com/rpcxio/libkv/store"
	etcdV3 "github.com/rpcxio/rpcx-etcd/client"
	"github.com/sirupsen/logrus"
	"github.com/smallnest/rpcx/client"
)

var LogicRpcClient client.XClient
var once sync.Once

type RpcLogic struct {
}

var RpcLogicObj *RpcLogic

func InitLogicRpcClient() {
	once.Do(func() {
		etcdConfigOption := &store.Config{
			ClientTLS:         nil,
			TLS:               nil,
			ConnectionTimeout: time.Duration(config.Conf.Common.CommonEtcd.ConnectionTimeout) * time.Second,
			Bucket:            "",
			PersistConnection: true,
			Username:          config.Conf.Common.CommonEtcd.UserName,
			Password:          config.Conf.Common.CommonEtcd.Password,
		}
		d, err := etcdV3.NewEtcdV3Discovery(
			config.Conf.Common.CommonEtcd.BasePath,
			config.Conf.Common.CommonEtcd.ServerPathLogic,
			[]string{config.Conf.Common.CommonEtcd.Host},
			true,
			etcdConfigOption,
		)
		if err != nil {
			err = config.Retry(`[api]init connect rpc etcd discovery client`, func() (bool, error) {
				d, err = etcdV3.NewEtcdV3Discovery(
					config.Conf.Common.CommonEtcd.BasePath,
					config.Conf.Common.CommonEtcd.ServerPathLogic,
					[]string{config.Conf.Common.CommonEtcd.Host},
					true,
					etcdConfigOption,
				)
				return err != nil, err
			}, err, 10)
			if err != nil {
				logrus.Fatalf("init connect rpc etcd discovery client fail: %s", err.Error())
			}
		}
		LogicRpcClient = client.NewXClient(config.Conf.Common.CommonEtcd.ServerPathLogic, client.Failtry, client.RandomSelect, d, client.DefaultOption)
		RpcLogicObj = new(RpcLogic)
	})
	if LogicRpcClient == nil {
		logrus.Fatalf("get logic rpc client nil")
	}
}

func (rpc *RpcLogic) Login(req *proto.LoginRequest) (code int, authToken string, msg string) {
	reply := &proto.LoginResponse{}
	err := LogicRpcClient.Call(context.Background(), "Login", req, reply)
	if err != nil {
		msg = err.Error()
	}
	code = reply.Code
	authToken = reply.AuthToken
	return
}

func (rpc *RpcLogic) Register(req *proto.RegisterRequest) (code int, authToken string, msg string) {
	reply := &proto.RegisterReply{}
	err := LogicRpcClient.Call(context.Background(), "Register", req, reply)
	if err != nil {
		msg = err.Error()
	}
	code = reply.Code
	authToken = reply.AuthToken
	return
}

func (rpc *RpcLogic) GetUserNameByUserId(req *proto.GetUserInfoRequest) (code int, userName string) {
	reply := &proto.GetUserInfoResponse{}
	err := LogicRpcClient.Call(context.Background(), "GetUserInfoByUserId", req, reply)
	if err != nil {
		logrus.Error(err)
	}
	code = reply.Code
	userName = reply.UserName
	return
}

func (rpc *RpcLogic) CheckAuth(req *proto.CheckAuthRequest) (code int, userId int, userName string) {
	reply := &proto.CheckAuthResponse{}
	err := LogicRpcClient.Call(context.Background(), "CheckAuth", req, reply)
	if err != nil {
		logrus.Error(err)
	}
	code = reply.Code
	userId = reply.UserId
	userName = reply.UserName
	return
}

func (rpc *RpcLogic) Logout(req *proto.LogoutRequest) (code int) {
	reply := &proto.LogoutResponse{}
	err := LogicRpcClient.Call(context.Background(), "Logout", req, reply)
	if err != nil {
		logrus.Error(err.Error())
	}
	code = reply.Code
	return
}

func (rpc *RpcLogic) Push(req *proto.Send) (code int, msg string) {
	reply := &proto.SuccessReply{}
	err := LogicRpcClient.Call(context.Background(), "Push", req, reply)
	if err != nil {
		msg = err.Error()
		return
	}
	code = reply.Code
	msg = reply.Msg
	return
}

func (rpc *RpcLogic) PushRoom(req *proto.Send) (code int, msg string) {
	reply := &proto.SuccessReply{}
	err := LogicRpcClient.Call(context.Background(), "PushRoom", req, reply)
	if err != nil {
		msg = err.Error()
		return
	}
	code = reply.Code
	msg = reply.Msg
	return
}

func (rpc *RpcLogic) Count(req *proto.Send) (code int, msg string) {
	reply := &proto.SuccessReply{}
	err := LogicRpcClient.Call(context.Background(), "Count", req, reply)
	if err != nil {
		msg = err.Error()
		return
	}
	code = reply.Code
	msg = reply.Msg
	return
}

func (rpc *RpcLogic) GetRoomInfo(req *proto.Send) (code int, msg string) {
	reply := &proto.SuccessReply{}
	err := LogicRpcClient.Call(context.Background(), "GetRoomInfo", req, reply)
	if err != nil {
		msg = err.Error()
		return
	}
	code = reply.Code
	msg = reply.Msg
	return
}
