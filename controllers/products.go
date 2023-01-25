package controllers

import (
	"html/template"
	"net/http"
	"strconv"

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

func AddProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "AddProduct", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		convertPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err.Error())
		}

		convertQuantity, err := strconv.Atoi(quantity)

		if err != nil {
			panic(err.Error())
		}

		models.InsertProduct(name, description, convertPrice, convertQuantity)
	}

	http.Redirect(w, r, "/", 301)
}
