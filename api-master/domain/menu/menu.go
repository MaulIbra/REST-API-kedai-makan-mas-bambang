package menu

import "github.com/maulIbra/clean-architecture-go/api-master/domain/category"

type Menu struct {
	MenuiD string `json:"menuId"`
	Category category.Category `json:"jenis"`
	MenuName string `json:"menuName"`
	Stock int `json:"stock"`
	Price int `json:"price"`
	MenuActive int `json:"menuActive"`
}
