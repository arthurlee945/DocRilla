package main

import (
	"context"
	"fmt"
	"log"

	"github.com/arthurlee945/Docrilla/internal/config"

	"github.com/arthurlee945/Docrilla/internal/db"
)

// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/

func main() {
	cfg, err := config.Load(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	dbConn, err := db.Connect(cfg.DSN)
	if err != nil {
		log.Fatalln(err)
	}

	defer dbConn.Close()

	if err := db.DropAllTable(dbConn); err != nil {
		fmt.Println(err)
	}
	if err := db.InitializeTable(dbConn); err != nil {
		fmt.Println(err)
	}
	if err := db.Seed(dbConn); err != nil {
		fmt.Println(err)
	}
}
