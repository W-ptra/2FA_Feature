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
	database.Migration()

	addr := fmt.Sprintf("%v:%v",os.Getenv("HOST"),os.Getenv("PORT"))
	server := NewServer(addr)
	server.run()
}