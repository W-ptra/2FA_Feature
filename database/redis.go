package database

import (
	"context"
	"log"
	"github.com/redis/go-redis/v9"
)

func GetRedisConnection()(*redis.Client,error){
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	ctx := context.Background()
	pong,err := rdb.Ping(ctx).Result()
	if err != nil{
		return nil,err
	}
	log.Println("connected to redis: ",pong)
	return rdb,nil
}