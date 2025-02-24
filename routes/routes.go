package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfajarama/go-struct-pos/controllers"
)

func SetupRoutes(app *fiber.App, controller *controllers.ProductController) {
	productGroup := app.Group("/products")
	productGroup.Post("/", controller.CreateProduct)
	productGroup.Get("/", controller.GetAllProducts)
	productGroup.Get("/:product_id", controller.GetProductByID)
	productGroup.Put("/:product_id", controller.UpdateProduct)
	productGroup.Delete("/:product_id", controller.DeleteProduct)
}
