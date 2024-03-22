package server

import "net/http"

type App struct {
	mux *http.ServeMux
}

func NewApp() *App {
	return &App{
		mux: routeInitializer(),
	}
}

func (a *App) Run(port string) error {
	return nil
}

func routeInitializer() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
