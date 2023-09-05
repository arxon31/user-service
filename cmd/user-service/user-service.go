package main

import (
	"github.com/arxon31/user-service/internal/user"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug}))
	logger.Info("Starting user-service")
	logger.Info("create router")
	router := chi.NewRouter()
	handler := user.NewHandler(logger)
	handler.Register(router)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}
	if err := server.Serve(listener); err != nil {
		logger.Error("%v", err)
		os.Exit(5)
	}

}
