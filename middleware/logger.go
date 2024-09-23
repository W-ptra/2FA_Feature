package middleware

import (
	"net/http"
	"log"
)

func Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){
		log.Printf("%v %v %v\n",r.Method,r.Host,r.URL.Path)
		next.ServeHTTP(w,r)
	}
}