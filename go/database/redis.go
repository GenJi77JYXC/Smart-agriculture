package database

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)

var rdb *redis.Client

func InitRedis() {
	db := viper.GetInt("redis.db")
	addr := viper.GetString("redis.addr")
	password := viper.GetString("redis.password")
	rdb = newClient(db, addr, password)
	fmt.Println("redis初始完成")
}

func newClient(db int, addr, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return client
}

func GetRDB() *redis.Client {
	return rdb
}

func RedisSetKey(key, val string) {
	// 将该条token失效的时间设置为token失效的最大时间 24 * time.Hour * 7
	err := rdb.Set(key, val, 24*time.Hour*7).Err()
	if err != nil {
		fmt.Println("RedisSetKey出错")
		panic(err)
	}
}

func RedisGetKey(key string) string {
	val, err := rdb.Get(key).Result()
	if err != nil {
		fmt.Println("RedisGetKey出错:", err)
	}
	// 如果val等于空字符串， 证明这个key没有被放入redis， 也就是这个token没被注销
	return val
}
