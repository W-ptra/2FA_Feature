package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"github.com/W-ptra/2FA-Feature/database"
	"github.com/W-ptra/2FA-Feature/service"
)

type LoginUser struct{
	Email	string 		`json:"email"`
	Password string 	`json:"password"`
}

type RegisterUser struct{
	Name	string 			`json:"name"`
	Email	string 			`json:"email"`
	Password string 		`json:"password"`
	ConfirmPassword	string 	`json:"confirmPassword"`
}

type Otp struct{
	Email 	string			`json:"email"`
	Code	string			`json:"code"`
}

type Message struct{
	Message string
}

func setRespond(w http.ResponseWriter,r *http.Request,message interface{},status int){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}

func PostLogin(w http.ResponseWriter,r *http.Request){
	body,_ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var user LoginUser
	if err := json.Unmarshal(body,&user);err!=nil{
		setRespond(w,r,Message{"Can't unmarshal to json"},400)
		return
	}

	if user.Email == "" || user.Password == ""{
		setRespond(w,r,Message{"field email or password is undefined or empty"},400)
		return
	}

	log.Println(user)

	userDB,err := database.GetUserByEmail(user.Email)
	if err!=nil && err.Error() == "record not found"{
		setRespond(w,r,Message{"User not found"},404)
		return
	}

	isMatch := service.ComparePassword(userDB.Password,user.Password)
	if !isMatch{
		setRespond(w,r,Message{"Wrong Password"},401)
		return
	}
	
	randomNumber := service.CreateOTPNumber()
	err = service.SendEmail(user.Email,randomNumber)
	
	if err != nil{
		setRespond(w,r,Message{"Wrong Password, cant send otp"},500)
		return
	}
	
	isOTPExist,_ := database.GetOTP(user.Email)
	
	if isOTPExist!=""{
		setRespond(w,r,Message{"OTP Already been sent, check your email"},200)
		return
	}

	err = database.SetOTP(user.Email,strconv.Itoa(randomNumber))
	if err != nil{
		setRespond(w,r,Message{"something went wrong, cant set otp"},500)
		return
	}

	setRespond(w,r,Message{"OTP has been sent, please chek your email"},200)
}

func PostRegister(w http.ResponseWriter,r *http.Request){
	body,_ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var user RegisterUser
	if err := json.Unmarshal(body,&user);err!=nil{
		setRespond(w,r,Message{"cant unmarsal to json"},400)
		return
	}

	if  user.Name == "" || user.Email == "" || user.Password == "" || user.ConfirmPassword == ""{
		setRespond(w,r,Message{"field name,email,password or confirmPassword is undefined or empty"},400)
		return
	}

	if user.Password != user.ConfirmPassword{
		setRespond(w,r,Message{"password and confirm password doesn't match"},400)
		return
	}

	log.Println(user)
	hashedPassword := service.HashPassword(user.Password)
	newUser := database.User{
		Name: user.Name,
		Email: user.Email,
		Password: hashedPassword,
	}
	
	err := database.CreateNewUser(newUser)
	if err!=nil{
		setRespond(w,r,Message{"something went wrong, can't create new user"},500)
		return
	}

	http.Redirect(w,r,"/login",http.StatusSeeOther)
}

func PostOtp(w http.ResponseWriter,r *http.Request){
	body,_ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var otp Otp
	if err:=json.Unmarshal(body,&otp); err!=nil{
		setRespond(w,r,Message{"can't unmarshal to json"},400)
		return
	}

	if otp.Code == ""{
		setRespond(w,r,Message{"otp code missing/empty"},400)
		return
	}
	log.Println(otp)
	rdsOTP,err := database.GetOTP(otp.Email)
	if err != nil{
		log.Println(err)
		setRespond(w,r,Message{"something went wrong"},500)
		return
	}
	
	if rdsOTP == ""{
		setRespond(w,r,Message{"No OTP been issue with corresponding email, or might expired"},409)
		return
	}

	if otp.Code != rdsOTP{
		setRespond(w,r,Message{"OTP Code is incorrect"},401)
		return
	}

	setRespond(w,r,Message{"OTP Code is Correct,Welcome"},200)
}