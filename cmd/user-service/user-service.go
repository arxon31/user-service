package main

import (
	"github.com/arxon31/user-service/internal/user"
	"github.com/go-chi/chi/v5"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	log.Println("Starting user-service")

	log.Println("create router")
	router := chi.NewRouter()
	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *chi.Mux) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}
	log.Fatalln(server.Serve(listener))
}
