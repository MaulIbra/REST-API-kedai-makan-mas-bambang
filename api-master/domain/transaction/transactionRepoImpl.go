package transaction

import (
	"database/sql"
	guuid "github.com/google/uuid"
	"github.com/maulIbra/clean-architecture-go/api-master/models"
	"github.com/maulIbra/clean-architecture-go/utils"
	"log"
)

type TransactionRepo struct {
	db *sql.DB
}


func NewTransactionRepo(db *sql.DB) ITransactionRepo{
	return &TransactionRepo{
		db: db,
	}
}

func (t TransactionRepo) GetTransaction(counter string) ([]*models.TransactionResponseTemp, error) {
	transactionTemp := []*models.TransactionResponseTemp{}

	stmt, err := t.db.Prepare(utils.SELECT_TRANSACTION)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(counter)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := models.TransactionResponseTemp{}
		err := rows.Scan(&p.TransactionDate, &p.TransactionId, &p.Menu.MenuId, &p.Menu.MenuName, &p.Menu.Quantity, &p.Menu.MenuPrice, &p.Menu.TotalPrice)
		if err != nil {
			return nil, err
		}
		transactionTemp = append(transactionTemp, &p)
	}

	return transactionTemp,nil
}

func (t TransactionRepo) GetTransactionByID(id string) ([]models.TransactionResponseTemp, error) {
	transactionTemp := []models.TransactionResponseTemp{}

	stmt, err := t.db.Prepare(utils.SELECT_TRANSACTION_BY_ID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := models.TransactionResponseTemp{}
		err := rows.Scan(&p.TransactionDate,&p.TransactionId,&p.Menu.MenuId,&p.Menu.MenuName,&p.Menu.Quantity,&p.Menu.MenuPrice,&p.Menu.TotalPrice)
		if err != nil {
			return nil, err
		}
		transactionTemp = append(transactionTemp, p)
	}
	return transactionTemp,nil
}

func (t TransactionRepo) PostTransaction(transaction *models.Transaction,updateStock map[string]int) error {
	id := guuid.New()
	transaction.TransactionId = id.String()
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.INSERT_TRANSACTION)
	defer stmt.Close()
	if err != nil {
		log.Print(err)
		tx.Rollback()
		return err
	}
	for _,val := range transaction.ListMenu{
		_, err = stmt.Exec(id,val.MenuId,transaction.TransactionDate,val.Quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	stmt,err = tx.Prepare(utils.UPDATE_STOCK_MENU)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	for val2 := range updateStock{
		_,err = stmt.Exec(updateStock[val2],val2)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}


func (t TransactionRepo) CheckMenuStock(id string) (*models.Menu, error) {
	var menu models.Menu
	stmt, err := t.db.Prepare(utils.SELECT_STOCK_MENU)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&menu.Stock,&menu.MenuName)
	if err != nil {
		return nil, err
	}
	return &menu, nil
}