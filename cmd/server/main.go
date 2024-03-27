package main

import (
	"fmt"
	"log"

	"github.com/arthurlee945/Docrilla/config"
)

func init() {
	if err := config.Initialize(".env"); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	config := config.New()
	fmt.Println("Init push", config.DatabaseUrl)
}
