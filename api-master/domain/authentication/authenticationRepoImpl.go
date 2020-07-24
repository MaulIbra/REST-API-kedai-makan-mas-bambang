package authentication

import (
	"database/sql"
	"errors"
	guuid "github.com/google/uuid"
	"github.com/maulIbra/clean-architecture-go/api-master/models"
	"github.com/maulIbra/clean-architecture-go/utils"
)

type AuthenticationRepo struct {
	db *sql.DB
}

func NewAuthenticationRepo(db *sql.DB) IAuthenticationRepo{
	return &AuthenticationRepo{
		db: db,
	}
}

func (a AuthenticationRepo) AddUserProfile(profile *models.Profile) error {
	id := guuid.New()
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.INSERT_USER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(id, profile.User.Username, profile.User.Password); err != nil {
		tx.Rollback()
		return err
	}

	stmt, err = tx.Prepare(utils.INSERT_PROFILE)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	profileID := guuid.New()
	if _, err := stmt.Exec(profileID, id, profile.NamaLengkap, profile.JenisKelamin, profile.Alamat); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (a AuthenticationRepo) ReadUserByEmail(username string) (*models.User, error) {
	stmt, err := a.db.Prepare(utils.SELECT_USER_BY_EMAIL)
	u := models.User{}
	if err != nil {
		return &u, err
	}
	errQuery := stmt.QueryRow(username).Scan(&u.UserID, &u.Username, &u.Password)

	if errQuery != nil {
		return &u, err
	}

	defer stmt.Close()
	return &u, nil
}

func (a AuthenticationRepo) ReadUser() (*[]models.Profile, error) {
	stmt, err := a.db.Prepare(utils.SELECT_USER)
	listProfile := []models.Profile{}
	if err != nil {
		return &listProfile, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return &listProfile, err
	}
	for rows.Next() {
		p := models.Profile{}
		err := rows.Scan(&p.User.UserID, &p.User.Username, &p.NamaLengkap, &p.JenisKelamin, &p.Alamat)
		if err != nil {
			return &listProfile, err
		}
		listProfile = append(listProfile, p)
	}

	return &listProfile, nil
}

func (a AuthenticationRepo) DeleteUser(userID string) error {
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.DELETE_USER_PROFILE)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err := stmt.Exec(userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("gagal delete, user id tidak di temukan")
	}

	stmt, err = tx.Prepare(utils.DELETE_USER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	if _, err := stmt.Exec(userID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
