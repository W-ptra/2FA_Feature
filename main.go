package main

import (
	"fmt"
	"github.com/W-ptra/2FA-Feature/database"
)

func main(){
	db,err := database.GetConnection()
	if err!=nil{
		fmt.Println("error connecting database",err)
	}
	err = db.AutoMigrate(&database.User{})
	if err!=nil{
		fmt.Println("error migrating database",err)
	}

	rds,errorr := database.GetRedisConnection()
	fmt.Println(rds)
	fmt.Println(errorr)

	server := NewServer("0.0.0.0:8001")
	server.run()
}