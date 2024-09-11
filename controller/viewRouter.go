package controller

import (
	"net/http"
)

func GetLogin(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"./public/login.html")
}

func GetRegister(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"./public/register.html")
}