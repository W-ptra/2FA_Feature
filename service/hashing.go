package service

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainPassword string)string{
	hashedPassword,_ := bcrypt.GenerateFromPassword([]byte(plainPassword),bcrypt.DefaultCost)
	return string(hashedPassword)
}

func comparePassword(hashedPassword string,plainPassword string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(plainPassword))
	if err != nil{
		return false //password not match
	}
	return true //password match
}