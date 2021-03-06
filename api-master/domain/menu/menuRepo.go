package menu

import "github.com/maulIbra/clean-architecture-go/api-master/models"

type IMenuRepo interface {
	GetMenu(offset,lengthRow int) ([]*models.Menu, error)
	GetCountMenu() (*int,error)
	GetMenuByID(id string) (*models.Menu, error)
	PostMenu(menu *models.Menu) error
	UpdateMenu(menu *models.Menu) error
	DeleteMenu(id string) error
}
