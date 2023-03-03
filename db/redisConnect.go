package db

import (
	"MiniDouyin/config"
	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool

func InitRdb() {
	Pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			dial, err := redis.Dial(config.RdbNetwork, config.RdbAddress, redis.DialDatabase(config.RdbUseDatabase), redis.DialPassword(config.RdbPassword))
			if err != nil {
				panic(err)
			}
			return dial, nil
		}}
}
