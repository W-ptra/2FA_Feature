package database

import (
	"fmt"
)

func Migration(){
	db,err := GetConnection()
	if err!=nil{
		fmt.Println("error connecting to database",err)
	}
	err = db.AutoMigrate(&User{})
	if err!=nil{
		fmt.Println("error migrating database",err)
	}
	fmt.Println("postgres migration/test success")

	rds,err := GetRedisConnection()
	if err!=nil{
		fmt.Println("error connecting to redis",err)
	}
	fmt.Println("redis test connection success: ",rds)
}