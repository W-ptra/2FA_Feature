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
	_,err := rdb.Ping(ctx).Result()
	if err != nil{
		return nil,err
	}
	return rdb,nil
}

func SetOTP(email,otp string)error{
	rdb,err := GetRedisConnection()
	if err != nil {
        return err
    }
	ctx := context.Background()
	err = rdb.Set(ctx,email, otp, 4*time.Minute).Err()
    if err != nil {
        return err
    }
	log.Printf("Redis: save OTP %v with key %v\n",otp,email);
	return nil
}

func GetOTP(email string)(string,error){
	rdb,err := GetRedisConnection()
	if err != nil {
        return "",err
    }
	ctx := context.Background()
	otp,err := rdb.Get(ctx,email).Result()
	if err == redis.Nil {
        return "",err
    } else if err != nil {
		return "",err
	}
	return otp,nil
}