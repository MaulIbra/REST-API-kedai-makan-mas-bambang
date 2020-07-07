package menu

import "errors"

type menuUsecase struct {
	menuRepo IMenuRepo
}

func NewMenuUsecase(repo IMenuRepo) IMenuUsecase{
	return &menuUsecase{
		menuRepo: repo,
	}
}


func (m menuUsecase) GetMenu() ([]*Menu, error) {
	listMenu,err := m.menuRepo.GetMenu()
	if err != nil {
		return nil, err
	}
	return listMenu,nil
}

func (m menuUsecase) GetMenuByID(id string) (*Menu, error) {
	if len(id) <= 0 {
		return nil, errors.New("no params")
	}
	menu,err := m.menuRepo.GetMenuByID(id)
	if err != nil {
		return nil, err
	}
	return menu,nil
}

func (m menuUsecase) PostMenu(menu *Menu) error {
	err := m.menuRepo.PostMenu(menu)
	if err != nil {
		return err
	}
	return nil
}

func (m menuUsecase) UpdateMenu(menu *Menu) error {
	err := m.menuRepo.UpdateMenu(menu)
	if err != nil {
		return err
	}
	return nil
}

func (m menuUsecase) DeleteMenu(id string) error {
	err := m.menuRepo.DeleteMenu(id)
	if err != nil {
		return err
	}
	return nil
}