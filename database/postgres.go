package database
import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func GetConnection()(*gorm.DB,error){
	err := godotenv.Load()
	if err!=nil{
		fmt.Println("error loading environment variable",err)
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Shanghai",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_DB"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_SSLMODE"),
	)
	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err!=nil{
		return nil,err
	}

	return db,nil
}

func CreateNewUser(db *gorm.DB,newUser User)error{
	operation := db.Create(&newUser)
	return operation.Error
}

func GetUserByEmail(db *gorm.DB,email string)(User,error){
	var user User
	operation := db.First(&user,"Email = ?",email)
	return user,operation.Error
}