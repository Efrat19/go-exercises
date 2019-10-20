package router

import (
	. "github.com/Efrat19/gophercises/cyoa/controllers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Home",
		"GET",
		"/",
		Home,
	},
	Route{
		"Read",
		"GET",
		"/read/{chapter}",
		Read,
	},
}
