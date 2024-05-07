package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/arthurlee945/Docrilla/internal/config"
	"github.com/arthurlee945/Docrilla/internal/logger"
	"github.com/arthurlee945/Docrilla/internal/server"
	"github.com/arthurlee945/Docrilla/internal/service/field"
	"github.com/arthurlee945/Docrilla/internal/service/project"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

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

func run(ctx context.Context, w io.Writer, _ []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	srvLogger := logger.New()
	cfg, err := config.Load(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	dbConn, err := db.Connect(cfg.DSN)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			srvLogger.Error("db close error", zap.Error(err))
		}
		cancel()
	}()
	// Might not need this
	addr := flag.String("http", fmt.Sprintf(":%s", cfg.Port), "HTTP service address (e.g.. '127.0.0.1:8080' or ':8080')")
	flag.Parse()

	// proj
	projService := project.NewService(project.NewRepository(dbConn), field.NewRepository(dbConn))

	srv := server.New(ctx, projService)
	httpServer := http.Server{
		Addr:    *addr,
		Handler: srv,
	}
	go func() {
		fmt.Fprintln(w, "Listening on "+httpLink(*addr, false))
		if err := httpServer.ListenAndServe(); err != nil && err == http.ErrServerClosed {
			srvLogger.Error("ListenAndServe failed", zap.Error(err))
			os.Exit(1)
		}
	}()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			srvLogger.Error("failed shutting down gracefully", zap.Error(err))
		}
	}()
	wg.Wait()
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

func httpLink(addr string, secure bool) string {
	if addr[0] == ':' {
		addr = "localhost" + addr
	}
	if secure {
		return "https://" + addr
	}
	return "http://" + addr
}
