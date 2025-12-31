package main

import (
	"fmt"
	"net/http"

	"github.com/Akshat-Kumar-work/golang-rest-api/internal/config"
	"go.uber.org/zap"
)

func main() {

	//load config file
	cfg := config.LoadConfig()

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to api"))
		fmt.Println("welcome to api")
	})

	var logger *zap.Logger

	if cfg.Env == "prod" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
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

	//start sever
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("err is", err)
	}
}
