package database

import (
	"context"
	"log"
	"time"
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

func SetOTP(rdb *redis.Client,email,otp string)error{
	ctx := context.Background()
	err := rdb.Set(ctx,email, otp, 2*time.Minute).Err()
    if err != nil {
        return err
    }
	return nil
}

func CheckOTP(rdb *redis.Client,email string)(string,error){
	ctx := context.Background()
	otp,err := rdb.Get(ctx,email).Result()
	if err != nil {
        return "",err
    }
	return otp,nil
}