package api

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type HttpHandler func(wr http.ResponseWriter, req *http.Request) HttpError

func NewRouter() http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	//recovery middleware for any panics in the handlers
	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	//add middleware for all routes
	n := negroni.New(recovery)
	//add some top level routes

	r.HandleFunc("/sys/info/health", RouteErrorHandler(HealthHandler))
	r.HandleFunc("/sys/info/ping", RouteErrorHandler(Ping))
	r.HandleFunc("/docker/images", RouteErrorHandler(ImagePull)).Methods("POST")
	//wire up middleware and router
	n.UseHandler(r)

	return n //negroni implements the http.Handler interface
}
