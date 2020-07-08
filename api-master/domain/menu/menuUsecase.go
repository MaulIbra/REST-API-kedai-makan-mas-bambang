package menu

import "github.com/maulIbra/clean-architecture-go/api-master/models"

type IMenuUsecase interface {
	GetMenu() ([]*models.Menu, error)
	GetMenuByID(id string) (*models.Menu, error)
	PostMenu(menu *models.Menu) error
	UpdateMenu(menu *models.Menu) error
	DeleteMenu(id string) error
}
