package models

import (
	"github.com/Porto-08/app-web-crud/config/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := db.ConectWithDatabase()

	selectProducts, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}

	products := []Product{}

	for selectProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer db.Close()

	return products
}

func InsertProduct(name string, description string, price float64, quantity int) {
	db := db.ConectWithDatabase()

	insertProduct, err := db.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProduct(productId string) {
	db := db.ConectWithDatabase()

	deleteProduct, err := db.Prepare("delete from products where id = $1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(productId)

	defer db.Close()
}
