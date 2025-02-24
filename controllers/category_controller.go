package controllers

import (
	"github.com/tfajarama/go-struct-pos/models"
	"github.com/tfajarama/go-struct-pos/repositories"
)

type CategoryController struct {
	repo repositories.CategoryRepository
}

func (controller *CategoryController) Create(category models.Category) (models.Category, error) {
	return controller.repo.Create(category)
}

func (controller *CategoryController) FindById(id uint64) (models.Category, error) {
	return controller.repo.FindById(id)
}

func (controller *CategoryController) FindAll() ([]models.Category, error) {
	return controller.repo.FindAll()
}

func (controller *CategoryController) Update(category models.Category) (models.Category, error) {
	return controller.repo.Update(category)
}
