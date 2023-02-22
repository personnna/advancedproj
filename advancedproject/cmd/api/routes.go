package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/justchecking", app.justcheckingHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tickets", app.createTicketHandler)
	router.HandlerFunc(http.MethodGet, "/v1/tickets/:id", app.showTicketHandler)
	return router
}
