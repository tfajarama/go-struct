package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfajarama/go-struct-pos/models"
	"github.com/tfajarama/go-struct-pos/services"
)

type ProductController struct {
	productService services.ProductService
}

func (c *ProductController) CreateProduct(ctx *fiber.Ctx) error {
	var product models.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.productService.CreateProduct(&product); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(product)
}

func (c *ProductController) GetAllProducts(ctx *fiber.Ctx) error {
	products, err := c.productService.GetAllProducts()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(200).JSON(products)
}

func (c *ProductController) GetProductByID(ctx *fiber.Ctx) error {
	id := ctx.Params("product_id")
	product, err := c.productService.GetProductByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	return ctx.Status(200).JSON(product)
}

func (c *ProductController) UpdateProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("product_id")
	var product models.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.productService.UpdateProduct(id, &product); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(200).JSON(product)
}

func (c *ProductController) DeleteProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("product_id")
	if err := c.productService.DeleteProduct(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService}
}
