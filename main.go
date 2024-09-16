package main

import (
	"fmt"
	"github.com/W-ptra/2FA-Feature/database"
)

func main(){
	db,err := database.GetConnection()
	if err!=nil{
		fmt.Println("error connecting to database",err)
	}
	err = db.AutoMigrate(&database.User{})
	if err!=nil{
		fmt.Println("error migrating database",err)
	}

	rds,err := database.GetRedisConnection()
	fmt.Println(rds)
	if err!=nil{
		fmt.Println("error connecting to redis",err)
	}

	server := NewServer("0.0.0.0:8001")
	server.run()
}