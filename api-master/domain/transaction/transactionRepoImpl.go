package transaction

import (
	"database/sql"
	guuid "github.com/google/uuid"
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

func (t TransactionRepo) GetTransaction(counter string) ([]*TransactionResponseTemp, error) {
	transactionTemp := []*TransactionResponseTemp{}

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
		p := TransactionResponseTemp{}
		err := rows.Scan(&p.TransactionDate, &p.TransactionId, &p.Menu.MenuId, &p.Menu.MenuName, &p.Menu.Quantity, &p.Menu.MenuPrice, &p.Menu.TotalPrice)
		if err != nil {
			return nil, err
		}
		transactionTemp = append(transactionTemp, &p)
	}

	return transactionTemp,nil
}

func (t TransactionRepo) GetTransactionByID(id string) ([]TransactionResponseTemp, error) {
	transactionTemp := []TransactionResponseTemp{}

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
		p := TransactionResponseTemp{}
		err := rows.Scan(&p.TransactionDate,&p.TransactionId,&p.Menu.MenuId,&p.Menu.MenuName,&p.Menu.Quantity,&p.Menu.MenuPrice,&p.Menu.TotalPrice)
		if err != nil {
			return nil, err
		}
		transactionTemp = append(transactionTemp, p)
	}
	return transactionTemp,nil
}

func (t TransactionRepo) PostTransaction(transaction *Transaction) error {
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
	return tx.Commit()
}

