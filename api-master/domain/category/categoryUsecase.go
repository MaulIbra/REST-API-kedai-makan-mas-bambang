package category

import "github.com/maulIbra/clean-architecture-go/api-master/models"

type ICategoryUsecase interface {
	GetCategory() ([]*models.Category, error)
}
