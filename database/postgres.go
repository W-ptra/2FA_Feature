package database
import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var postgresConnection *gorm.DB

func GetConnection()(*gorm.DB,error){
	if postgresConnection == nil{
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

		postgresConnection = db
	}

	return postgresConnection,nil
}

func CreateNewUser(newUser User)error{
	db,err := GetConnection()
	if err!=nil{
		return err
	}
	operation := db.Create(&newUser)
	return operation.Error
}

func GetUserByEmail(email string)(User,error){
	db,err := GetConnection()
	if err!=nil{
		return User{},err
	}
	var user User
	operation := db.First(&user,"Email = ?",email)
	return user,operation.Error
}