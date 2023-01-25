package controllers

import (
	"html/template"
	"net/http"

	"github.com/Porto-08/app-web-crud/models"
)

// config templates
var temp = template.Must(
	template.ParseGlob("templates/*.html"),
)

func Index(w http.ResponseWriter, r *http.Request) {
	AllProducts := models.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", AllProducts)
}
