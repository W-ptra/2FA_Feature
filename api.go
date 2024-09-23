package main

import (
	"log"
	"net/http"
	"github.com/W-ptra/2FA-Feature/controller"
)

type Server struct{
	addr string
}

func NewServer(addr string) *Server{
	return &Server{
		addr: addr,
	}
}

func (s *Server) run() error{
	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public"))
	
	router.Handle("/public/",fs)
	router.HandleFunc("/",controller.RedirectToLogin)	
	router.HandleFunc("GET /login",controller.GetLogin)	
	router.HandleFunc("POST /login",controller.PostLogin)
	router.HandleFunc("GET /register",controller.GetRegister)
	router.HandleFunc("POST /register",controller.PostRegister)
	router.HandleFunc("POST /otp",controller.PostOtp)

	server := http.Server{
		Addr: s.addr,
		Handler: router,
	}

	log.Printf("Listening to %v\n",s.addr)
	log.Printf("Open http://localhost:8001/login to proceed further\n")
	return server.ListenAndServe()
}