package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
	Code	string			`json:"Code"`
}

func PostLogin(w http.ResponseWriter,r *http.Request){
	body,_ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var user LoginUser
	if err := json.Unmarshal(body,&user);err!=nil{
		http.Error(w,"missing field: name or password",http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == ""{
		http.Error(w,"field email or password is undefined or empty",http.StatusBadRequest)
		return
	}

	db,err := database.GetConnection()
	if err!=nil{
		http.Error(w,"something went wrong",http.StatusInternalServerError)
		return
	}

	log.Println(user)

	userDB,errors := database.GetUserByEmail(db,user.Email)
	if errors!=nil{
		http.Error(w,"something went wrong",http.StatusInternalServerError)
		return
	}

	isMatch := service.ComparePassword(userDB.Password,user.Password)
	if !isMatch{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"Message":"Wrong Password"})
		return
	}

	randomNumber := service.CreateOTPNumber()
	err = service.SendEmail(user.Email,randomNumber)
	if err != nil{
		http.Error(w,"something went wrong",http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"Message":"Login Sucessfully","otp":randomNumber})
}

func PostRegister(w http.ResponseWriter,r *http.Request){
	body,_ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var user RegisterUser
	if err := json.Unmarshal(body,&user);err!=nil{
		http.Error(w,"Bad request",http.StatusBadRequest)
		return
	}

	if  user.Name == "" || user.Email == "" || user.Password == "" || user.ConfirmPassword == ""{
		http.Error(w,"field name,email,password or confirmPassword is undefined or empty",http.StatusBadRequest)
		return
	}

	if user.Password != user.ConfirmPassword{
		http.Error(w,"password and confirm password doesn't match",http.StatusBadRequest)
		return
	}

	log.Println(user)
	hashedPassword := service.HashPassword(user.Password)
	newUser := database.User{
		Name: user.Name,
		Email: user.Email,
		Password: hashedPassword,
	}
	db,err := database.GetConnection()
	if err!=nil{
		http.Error(w,"something went wrong",http.StatusInternalServerError)
		return
	}
	
	err = database.CreateNewUser(db,newUser)
	if err!=nil{
		http.Error(w,"something went wrong",http.StatusInternalServerError)
		return
	}

	http.Redirect(w,r,"/login",http.StatusSeeOther)
}

func PostOtp(w http.ResponseWriter,r *http.Request){
	body,_ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var otp Otp
	if err:=json.Unmarshal(body,&otp); err!=nil{
		http.Error(w,"Bad request",http.StatusBadRequest)
		return
	}

	if otp.Code == ""{
		http.Error(w,"Bad request",http.StatusBadRequest)
		return
	}

	//check otp code at redis

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"Message":otp.Code})
}