package store_test

import (
	"context"
	"log"
	"testing"

	"github.com/arthurlee945/Docrilla/internal/config"
	"github.com/arthurlee945/Docrilla/internal/db"
	"github.com/arthurlee945/Docrilla/internal/service/project/store"
)

func TestGetProjectOverview(t *testing.T) {
	projStore := getProjectStore()
}

func getProjectStore() *store.Store {
	cfg, err := config.Load(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	dbConn, err := db.Connect(cfg.DSN)
	if err != nil {
		log.Fatalln(err)
	}
	defer dbConn.Close()
	return store.NewStore(dbConn)
}
