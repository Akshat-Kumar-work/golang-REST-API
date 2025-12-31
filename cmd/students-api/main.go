package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Akshat-Kumar-work/golang-rest-api/internal/config"
	"github.com/Akshat-Kumar-work/golang-rest-api/pkg/logger"
	"go.uber.org/zap"
)

func main() {

	//load config file
	cfg := config.LoadConfig()

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to api"))
		logger.Info("home endpoint called")
	})

	err := logger.Initialize(logger.LogConfig(cfg.LogConfig))
	if err != nil {
		log.Fatal("logger not initialized", err)
	}

	defer logger.Sync() //write or flush log before main func exits ,as zap buffer log into memory to improve speed

	// Log server start info
	logger.Info("Starting server",
		zap.String("env", cfg.Env),
		zap.String("storage_path", cfg.StoragePath),
		zap.String("address", cfg.HTTPServer.Address),
	)

	//create server
	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}
	// Channel to listen for OS signals
	done := make(chan os.Signal, 1)

	//if any signal like interupt come from os notify to done channel
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server error", zap.Error(err))
		}
	}()

	// Wait for shutdown signal
	<-done
	logger.Info("shutting down server")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("server shutdown failed", zap.Error(err))
	}

	logger.Info("server exited gracefully")

}
