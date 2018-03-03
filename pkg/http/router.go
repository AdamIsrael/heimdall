package http

import (
	"github.com/adamisrael/heimdall/pkg/logger"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ListenAndServe(listen string, handler http.Handler) {
	log.Printf("Listening!")
	log.Fatal(http.ListenAndServe(
		listen,
		handler))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
