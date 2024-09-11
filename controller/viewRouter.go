package controller

import (
	"net/http"
)

func Root(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"./public/index.html")
}