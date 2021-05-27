package bk

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"log"
	"net/http"
)

type App struct {
	Router      *mux.Router
	Config      *Config
	Middlewares *Middleware
}

func (a *App) InitApp() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	log.Print("Init App start")
	a.Config = InitConfig()
	a.Router = mux.NewRouter()
	a.Middlewares = &Middleware{}
}
func (a *App) InitRouter() {
	m := alice.New(a.Middlewares.LoggingHandler, a.Middlewares.RecoverHandler)

	a.Router.Handle("/api/shorten", m.ThenFunc(a.createShortLink)).Methods("POST")
	a.Router.Handle("/api/info", m.ThenFunc(a.getShortInfo)).Methods("GET")
	a.Router.Handle("/api/{shortlink:[a-zA-Z0-9]{1,11}}", m.ThenFunc(a.redirect)).Methods("GET")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) createShortLink(w http.ResponseWriter, r *http.Request) {

}

func (a *App) getShortInfo(w http.ResponseWriter, r *http.Request) {

}
func (a *App) redirect(w http.ResponseWriter, r *http.Request) {

}
