package models

type Menu struct {
	MenuiD     string   `json:"menuId"`
	Category   Category `json:"jenis"`
	MenuName   string   `json:"menuName"`
	Stock      int      `json:"stock"`
	Price      int      `json:"price"`
	MenuActive int      `json:"menuActive"`
}
