/**
 * Created by lock
 * Date: 2019-09-22
 * Time: 22:37
 */
package db

import (
	"path/filepath"
	"sync"
	"time"

	"gochat/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

var dbMap = map[string]*gorm.DB{}
var syncLock sync.Mutex

func init() {
	initDB("gochat")
}

func initDB(dbName string) {
	var e error
	// if prod env , you should change mysql driver for yourself !!!
	realPath, _ := filepath.Abs("./")
	configFilePath := realPath + "/db/gochat.sqlite3"
	syncLock.Lock()
	dbMap[dbName], e = gorm.Open("sqlite3", configFilePath)
	dbMap[dbName].DB().SetMaxIdleConns(4)
	dbMap[dbName].DB().SetMaxOpenConns(20)
	dbMap[dbName].DB().SetConnMaxLifetime(8 * time.Second)
	if config.GetMode() == "dev" {
		dbMap[dbName].LogMode(true)
	}
	syncLock.Unlock()
	if e != nil {
		logrus.Errorf("connect db failed: %s", e.Error())
	}
}

func GetDb(dbName string) (db *gorm.DB) {
	db, _ = dbMap[dbName]
	return
}

type DbGoChat struct {
}

func (*DbGoChat) GetDbName() string {
	return "gochat"
}
