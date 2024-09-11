package main

import (
	"log"
	"net/http"
	"github.com/W-ptra/2FA-Feature/Controller"
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

	router.Handle("/public/",http.StripPrefix("/public/",fs))
	router.HandleFunc("GET /login",controller.GetLogin)	
	router.HandleFunc("GET /register",controller.GetRegister)

	server := http.Server{
		Addr: s.addr,
		Handler: router,
	}

	log.Printf("Server addr %v",s.addr)

	return server.ListenAndServe()
}