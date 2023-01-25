package main

import (
	"net/http"

	"github.com/Porto-08/app-web-crud/routes"
)

func main() {
	routes.LoadRoutes()

	http.ListenAndServe(":8000", nil)
}
