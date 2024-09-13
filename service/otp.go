package service

import(
	"math/rand"
)

func CreateOTPNumber()int{
	return rand.Intn(9000)+1000 // generate random 4 digit number
}