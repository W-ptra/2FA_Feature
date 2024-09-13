package service

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/resend/resend-go/v2"
)

func SendEmail(emailAddress string,OTP int)error{
	err := godotenv.Load()
	if err!=nil{
		return err
	}
	api_key := os.Getenv("RESEND_API_KEY")
	client := resend.NewClient(api_key)
	params := &resend.SendEmailRequest{
		From: "no-reply@wisnuputra.xyz",
		To: []string{emailAddress},
		Subject: "OTP",
		Text: fmt.Sprintf("Your OTP is %v and will be valid for 4 minutes",OTP),
	}
	send,err := client.Emails.Send(params)
	if err != nil{
		return err
	}
	fmt.Printf("Successfully send OTP to email address: %v, with send id %v\n",emailAddress,send.Id)
	return nil
}