package router
 
import (
	"net/http"
	. "github.com/Efrat19/gophercises/cyoa/controllers"
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