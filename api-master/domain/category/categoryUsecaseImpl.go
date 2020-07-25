package category

import "github.com/maulIbra/clean-architecture-go/api-master/models"

type categoryUsecaseImpl struct {
	repo ICategoryRepo
}

func NewCategoryUsecaseImpl(repo ICategoryRepo) ICategoryUsecase  {
	return &categoryUsecaseImpl{
		repo: repo,
	}
}

func (c categoryUsecaseImpl) GetCategory() ([]*models.Category, error) {
	categoryList,err := c.repo.GetCategory()
	if err != nil {
		return nil, err
	}
	return categoryList,nil
}
