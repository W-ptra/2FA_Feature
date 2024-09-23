package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var redisConnection *redis.Client

func GetRedisConnection()(*redis.Client,error){
	if redisConnection == nil {
		err := godotenv.Load()
		if err!=nil{
			log.Println("error loading environment variable",err)
		}
		rdb := redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%v:%v",os.Getenv("REDIS_HOST"),os.Getenv("REDIS_PORT")),
			Password: "",
			DB: 0,
		})
		ctx := context.Background()
		_,err = rdb.Ping(ctx).Result()
		if err != nil{
			return nil,err
		}
		redisConnection = rdb
	}
	return redisConnection,nil
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