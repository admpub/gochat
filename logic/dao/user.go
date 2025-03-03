/**
 * Created by lock
 * Date: 2019-09-22
 * Time: 22:53
 */
package dao

import (
	"context"
	"gochat/db"
	"gochat/logic/iface"
	"gochat/proto"
	"time"

	"github.com/pkg/errors"
)

func init() {
	iface.GetRepository = func(_ context.Context) iface.Repository {
		return new(User)
	}
}

var dbIns = db.GetDb("gochat")

type User struct {
	Id         int `gorm:"primary_key"`
	UserName   string
	Password   string
	CreateTime time.Time
	db.DbGoChat
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Register(req *proto.RegisterRequest) (userId int, err error) {
	u.UserName = req.Name
	u.Password = req.Password
	if u.UserName == "" || u.Password == "" {
		return 0, errors.New("user_name or password empty!")
	}
	oUser := u.CheckHaveUserName(u.UserName)
	if oUser.Id > 0 {
		return oUser.Id, nil
	}
	u.CreateTime = time.Now()
	if err = dbIns.Table(u.TableName()).Create(&u).Error; err != nil {
		return 0, err
	}
	return u.Id, nil
}

func (u *User) CheckHaveUserName(userName string) iface.UserData {
	var data User
	dbIns.Table(u.TableName()).Where("user_name=?", userName).Take(&data)
	return iface.UserData{
		Id:   data.Id,
		Name: data.UserName,
	}
}

func (u *User) Login(req *proto.LoginRequest) (iface.UserData, error) {
	userName := req.Name
	passWord := req.Password
	var data User
	var err error
	dbIns.Table(u.TableName()).Where("user_name=?", userName).Take(&data)
	if (data.Id == 0) || (passWord != data.Password) {
		err = errors.New("no this user or password error!")
	}
	return iface.UserData{
		Id:   data.Id,
		Name: data.UserName,
	}, err
}

func (u *User) GetUserNameByUserId(userId int) (userName string) {
	var data User
	dbIns.Table(u.TableName()).Where("id=?", userId).Take(&data)
	return data.UserName
}
