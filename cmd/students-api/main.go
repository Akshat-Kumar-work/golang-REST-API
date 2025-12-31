package main

import (
	"fmt"
	"log"

	"github.com/Akshat-Kumar-work/golang-rest-api/internal/config"
)

func main() {

	cfg, err := config.LoadConfig("local")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("students api")
	fmt.Println("Env", cfg.Env)
	fmt.Println("storage_path", cfg.StoragePath)
	fmt.Println("server_address ", cfg.HTTPServer.Address)
}
