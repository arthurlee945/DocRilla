package main

import (
	"context"
	"fmt"
	"log"

	"github.com/arthurlee945/Docrilla/config"
	"github.com/arthurlee945/Docrilla/internal/db"
	"github.com/arthurlee945/Docrilla/internal/model"
	projStore "github.com/arthurlee945/Docrilla/internal/project/store"
)

var cfg *config.Config

func init() {
	if err := config.Initialize(".env"); err != nil {
		log.Println("No .env file found")
	}
	cfg = config.New()
}

func main() {
	dbConn, err := db.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer dbConn.Close()

	//TESTING BLOCK
	projRepo := projStore.NewStore(dbConn)

	proj, projErr := projRepo.GetProjectDetail(context.Background(), &model.User{ID: 10}, "018ea1b1-b9ba-79af-81ce-81bae9930afa")
	if projErr != nil {
		fmt.Println("Query Errored")
		log.Fatalln(projErr)
	}
	fmt.Printf("%+v\n", proj)

}
