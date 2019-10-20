package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func NewRouter() *negroni.Negroni {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)

	return n
}
