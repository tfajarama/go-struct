package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfajarama/go-struct-pos/controllers"
)

func SetupRoutes(app *fiber.App, productController *controllers.ProductController, customerController *controllers.CustomerController) {
	productGroup := app.Group("/products")
	productGroup.Post("/", productController.CreateProduct)
	productGroup.Get("/", productController.GetAllProducts)
	productGroup.Get("/:product_id", productController.GetProductByID)
	productGroup.Put("/:product_id", productController.UpdateProduct)
	productGroup.Delete("/:product_id", productController.DeleteProduct)

	customerGroup := app.Group("/customers")
	customerGroup.Post("/", customerController.CreateCustomer)
	customerGroup.Get("/", customerController.GetAllCustomers)
	customerGroup.Get("/:customer_id", customerController.GetCustomerByID)
	customerGroup.Put("/:customer_id", customerController.UpdateCustomer)
	customerGroup.Delete("/:customer_id", customerController.DeleteCustomer)
}
