/**
 * Created by lock
 * Date: 2019-08-12
 * Time: 14:18
 */
package tools

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var RedisClientMap = map[string]*redis.Client{}
var syncLock sync.Mutex

type RedisOption struct {
	Address  string
	Password string
	Db       int
}

func GetRedisInstance(redisOpt RedisOption) *redis.Client {
	addr := redisOpt.Address
	db := redisOpt.Db
	password := redisOpt.Password
	syncLock.Lock()
	defer syncLock.Unlock()
	if redisCli, ok := RedisClientMap[addr]; ok {
		return redisCli
	}
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		Password:   password,
		DB:         db,
		MaxConnAge: 20 * time.Second,
	})
	RedisClientMap[addr] = client
	return RedisClientMap[addr]
}
