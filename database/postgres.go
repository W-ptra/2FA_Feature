package database
import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func GetConnection()(*gorm.DB,error){
	dsn := "host=localhost user=root password=root dbname=twofa_database port=8000 sslmode=disable TimeZone=Asia/Shanghai"
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