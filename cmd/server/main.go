package main

import (
	"fmt"
	"log"

	"github.com/arthurlee945/doc-rilla/config"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	config := config.New()
	fmt.Println("Init push", config.DatabaseUrl)
}
