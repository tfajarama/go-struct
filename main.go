package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tfajarama/go-struct-pos/controllers"
	"github.com/tfajarama/go-struct-pos/database"
	"github.com/tfajarama/go-struct-pos/repositories"
	"github.com/tfajarama/go-struct-pos/routes"
	"github.com/tfajarama/go-struct-pos/services"
)

func main() {
	app := fiber.New()

	database.ConnectDatabase()

	productRepo := repositories.NewProductRepository(database.DB)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	routes.SetupRoutes(app, productController)

	err := app.Listen(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
