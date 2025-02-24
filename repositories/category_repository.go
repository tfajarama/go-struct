package repositories

import "github.com/tfajarama/go-struct-pos/models"

type CategoryRepository interface {
	Create(category models.Category) (models.Category, error)
	FindById(id uint64) (models.Category, error)
	FindAll() ([]models.Category, error)
	Update(category models.Category) (models.Category, error)
}

type CategoryRepositoryImpl struct {
	CategoryRepository CategoryRepository
}

func (repo *CategoryRepositoryImpl) Create(category models.Category) (models.Category, error) {
	return repo.CategoryRepository.Create(category)
}

func (repo *CategoryRepositoryImpl) FindById(id uint64) (models.Category, error) {
	return repo.CategoryRepository.FindById(id)
}

func (repo *CategoryRepositoryImpl) FindAll() ([]models.Category, error) {
	return repo.CategoryRepository.FindAll()
}

func (repo *CategoryRepositoryImpl) Update(category models.Category) (models.Category, error) {
	return repo.CategoryRepository.Update(category)
}
