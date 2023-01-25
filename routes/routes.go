package routes

import (
	"net/http"

	"github.com/Porto-08/app-web-crud/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.AddProduct)
	http.HandleFunc("/insert", controllers.InsertProduct)
}
