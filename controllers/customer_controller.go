package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfajarama/go-struct-pos/models"
	"github.com/tfajarama/go-struct-pos/services"
)

type CustomerController struct {
	customerService services.CustomerService
}

func (c *CustomerController) CreateCustomer(ctx *fiber.Ctx) error {
	var customer models.Customer
	if err := ctx.BodyParser(&customer); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.customerService.CreateCustomer(&customer); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(customer)
}

func (c *CustomerController) GetAllCustomers(ctx *fiber.Ctx) error {
	customers, err := c.customerService.GetAllCustomers()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(200).JSON(customers)
}

func (c *CustomerController) GetCustomerByID(ctx *fiber.Ctx) error {
	id := ctx.Params("customer_id")
	customer, err := c.customerService.GetCustomerByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Customer not found"})
	}
	return ctx.Status(200).JSON(customer)
}

func (c *CustomerController) UpdateCustomer(ctx *fiber.Ctx) error {
	id := ctx.Params("customer_id")
	var customer models.Customer
	if err := ctx.BodyParser(&customer); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.customerService.UpdateCustomer(id, &customer); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(200).JSON(customer)
}

func (c *CustomerController) DeleteCustomer(ctx *fiber.Ctx) error {
	id := ctx.Params("customer_id")
	if err := c.customerService.DeleteCustomer(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}

func NewCustomerController(customerService services.CustomerService) *CustomerController {
	return &CustomerController{customerService}
}
