package menu

type IMenuRepo interface {
	GetMenu() ([]*Menu, error)
	GetMenuByID(id string) (*Menu, error)
	PostMenu(menu *Menu) error
	UpdateMenu(menu *Menu) error
	DeleteMenu(id string) error
}