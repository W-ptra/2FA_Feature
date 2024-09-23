package controller

import (
	"net/http"
)

func RedirectToLogin(w http.ResponseWriter,r *http.Request){
	http.Redirect(w,r,"/login",http.StatusSeeOther)
}

func GetLogin(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"./public/login.html")
}

func GetRegister(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"./public/register.html")
}
