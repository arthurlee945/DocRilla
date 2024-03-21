package server

import "net/http"

type App struct {
	httpServer *http.Server
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run(port string) error {
	return nil
}
