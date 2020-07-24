package authentication

import (
	"github.com/maulIbra/clean-architecture-go/api-master/models"
)

type IAuthenticationRepo interface {
	AddUserProfile(profile *models.Profile) error
	ReadUserByEmail(username string) (*models.User, error)
	ReadUser() (*[]models.Profile, error)
	DeleteUser(userID string) error
}
