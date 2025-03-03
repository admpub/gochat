package iface

import (
	"context"
	"gochat/proto"
)

type UserData struct {
	Id     int
	Name   string
	Avatar string
	Gender string
}

type Repository interface {
	Register(*proto.RegisterRequest) (userId int, err error)
	Login(*proto.LoginRequest) (data UserData, err error)
	CheckHaveUserName(userName string) (data UserData)
	GetUserNameByUserId(userId int) (userName string)
}

var GetRepository func(context.Context) Repository
