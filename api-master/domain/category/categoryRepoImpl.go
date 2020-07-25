package category

import (
	"database/sql"
	"github.com/maulIbra/clean-architecture-go/api-master/models"
	"github.com/maulIbra/clean-architecture-go/utils"
	"log"
)

type categoryRepoImpl struct {
	db *sql.DB
}

func NewCategoryRepoImpl(db *sql.DB) ICategoryRepo{
	return &categoryRepoImpl{
		db: db,
	}
}

func (c categoryRepoImpl) GetCategory() ([]*models.Category, error) {
	var categoryList []*models.Category
	stmt, err := c.db.Prepare(utils.SELECT_CATEGORY)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		category := models.Category{}
		err := rows.Scan(&category.CategoryID,&category.CategoryName)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		categoryList = append(categoryList, &category)
	}
	return categoryList, nil
}
