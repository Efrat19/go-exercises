package main

import (
	// "html/template"
	"net/http"
	// "encoding/json"
	"fmt"
	"github.com/Efrat19/gophercises/cyoa/router"
)


func main()  {

    r := router.NewRouter()

	fmt.Println("server is listening on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}




