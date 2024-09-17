package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/W-ptra/2FA-Feature/database"
)

func main(){
	err := godotenv.Load()
	if err!=nil{
		fmt.Println("error loading environment variable",err)
	}
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

	addr := fmt.Sprintf("%v:%v",os.Getenv("HOST"),os.Getenv("PORT"))
	server := NewServer(addr)
	server.run()
}