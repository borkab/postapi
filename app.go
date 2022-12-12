package app

import "github.com/gorilla/mux"

type App struct {
	Router *mux.Router
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.InitRoutes()
	return a
}

func (a *App) InitRoutes() {
	a.Router.HandleFunc("/", nil).Methods("GET")
}
