package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/arthurlee945/Docrilla/internal/config"
	"github.com/jmoiron/sqlx"

	"github.com/arthurlee945/Docrilla/internal/db"
)

// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
func main() {
	ctx := context.Background()

	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	cfg, err := config.Load(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	dbConn, err := db.Connect(cfg.DSN)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		dbConn.Close()
		cancel()
	}()
	http := flag.String("http", ":8080", "HTTP service address (e.g.. '127.0.0.1:8080' or ':8080')")
	flag.Parse()

	//TODO: add server here

	fmt.Fprintln(w, "Listening on "+httpLink(*http, false))
	return nil
}

func devtool(dbConn *sqlx.DB) error {
	if err := db.DropAllTable(dbConn); err != nil {
		return err
	}
	if err := db.InitializeTable(dbConn); err != nil {
		return err
	}
	if err := db.Seed(dbConn); err != nil {
		return err
	}
	return nil
}
