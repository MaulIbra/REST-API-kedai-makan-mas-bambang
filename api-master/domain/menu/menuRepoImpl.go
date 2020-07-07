package menu

import (
	"database/sql"
	"github.com/maulIbra/clean-architecture-go/utils"
	"log"
	guuid "github.com/google/uuid"
)

type menuRepo struct{
	db *sql.DB
}

func NewMenuRepo(db *sql.DB) IMenuRepo{
	return &menuRepo{db: db}
}
func (m *menuRepo) GetMenu() ([]*Menu, error) {
	menuList := []*Menu{}
	stmt, err := m.db.Prepare(utils.SELECT_MENU)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		menu := Menu{}
		err := rows.Scan(&menu.MenuiD,&menu.MenuName,&menu.Stock,&menu.Price,&menu.MenuActive,&menu.Category.CategoryID,&menu.Category.CategoryName)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		menuList = append(menuList, &menu)
	}
	return menuList, nil
}

func (m *menuRepo) GetMenuByID(id string) (*Menu, error) {
	menu := Menu{}
	stmt, err := m.db.Prepare(utils.SELECT_MENU_BY_ID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&menu.MenuiD,&menu.MenuName,&menu.Stock,&menu.Price,&menu.MenuActive,&menu.Category.CategoryID,&menu.Category.CategoryName)
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (m *menuRepo) PostMenu(menu *Menu) error {
	id := guuid.New()
	menu.MenuiD = id.String()
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.INSERT_MENU)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(id, menu.Category.CategoryID,menu.MenuName,menu.Stock,menu.MenuActive,menu.Price)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (m *menuRepo) UpdateMenu(menu *Menu) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_MENU)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(menu.Category.CategoryID,menu.MenuName,menu.Stock,menu.Price,menu.MenuActive,menu.MenuiD)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (m *menuRepo) DeleteMenu(id string) error {
	tx, err := m.db.Begin()
	idUpdate := 0
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.DELETE_MENU)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(idUpdate,id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
