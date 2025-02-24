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

	// Connect to the database
	database.ConnectDatabase()

	// Products Dependency Injection
	productRepo := repositories.NewProductRepository(database.DB)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	// Customers Dependency Injection
	customerRepo := repositories.NewCustomerRepository(database.DB)
	customerService := services.NewCustomerService(customerRepo)
	customerController := controllers.NewCustomerController(customerService)

	// Setup Routes
	routes.SetupRoutes(app, productController, customerController)

	// Start Server
	err := app.Listen(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
