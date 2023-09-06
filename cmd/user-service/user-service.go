package main

import (
	"fmt"
	"github.com/arxon31/user-service/internal/config"
	"github.com/arxon31/user-service/internal/logging"
	"github.com/arxon31/user-service/internal/user"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	cfg := config.InitConfig()
	logger := logging.GetLogger(cfg)
	logger.Info("Starting user-service")
	logger.Info("create router")
	router := chi.NewRouter()
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, logger, cfg)
}

func start(router *chi.Mux, logger *slog.Logger, config *config.Config) {
	var listener net.Listener
	var err error
	if config.Listen.Type == "port" {
		bindAddr := fmt.Sprintf("%s:%s", config.Listen.Host, config.Listen.Port)
		logger.Info("listen on tcp", "bind-address", bindAddr)
		listener, err = net.Listen("tcp", bindAddr)
		if err != nil {
			logger.Error("%v", err)
			os.Exit(5)
		}

	} else {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Error("%v", err)
			os.Exit(5)
		}
		logger.Info("create socket")

		socketPath := filepath.Join(appDir, "user-service.sock")
		logger.Debug("socket path:", "socket-path", socketPath)

		logger.Info("listen on unix socket", "socket-path", socketPath)
		listener, err = net.Listen("unix", socketPath)
		if err != nil {
			logger.Error("%v", err)
			os.Exit(5)
		}
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}
	if err = server.Serve(listener); err != nil {
		logger.Error("%v", err)
		os.Exit(5)
	}
}
