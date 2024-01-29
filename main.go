package main

import (
	"simple-ecommerce/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	//rute CRUD untuk product
	router.GET("/products", handlers.ListProducts(db))
	router.GET("/products/:id", handlers.GetProduct(db))
	router.POST("/products", handlers.CreateProduct(db))
	router.PUT("/products/:id", handlers.UpdateProduct(db))
	router.DELETE("/products/:id", handlers.DeleteProduct(db))

	router.Run(":8080")
}
