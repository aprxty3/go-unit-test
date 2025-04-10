package repositoy

import (
	"github.com/stretchr/testify/mock"
	"go-unit-test/entity"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repositoy *CategoryRepositoryMock) FindByID(id string) *entity.Category {
	args := repositoy.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}

	category := args.Get(0).(entity.Category)
	return &category

}
