package database
import "time"

type User struct{
	Id	int 			`gorm:"primaryKey:autoIncrement"`
	Name string
	Email string		`gorm:"index"`
	Password string
	CreatedAt time.Time
}