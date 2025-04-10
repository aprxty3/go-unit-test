package service

import (
	"errors"
	"go-unit-test/entity"
	"go-unit-test/repositoy"
)

type CategoryService struct {
	Repository repositoy.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindByID(id)
	if category == nil {
		return category, errors.New("Category Not Found")
	}
	return category, nil
}
