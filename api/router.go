package main

import (
	"github.com/gorilla/mux"
	"github.com/youkoulayley/serieall-api-go/api/middlewares"
	"net/http"
)

var router *mux.Router

func initializeRouter() *mux.Router {
	router = mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/api/v1.0")
	router.Use(middlewares.IsAuthenticate)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
