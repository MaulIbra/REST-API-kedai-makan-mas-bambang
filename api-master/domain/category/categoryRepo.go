package category

import "github.com/maulIbra/clean-architecture-go/api-master/models"

type ICategoryRepo interface {
	GetCategory() ([]*models.Category, error)
}

