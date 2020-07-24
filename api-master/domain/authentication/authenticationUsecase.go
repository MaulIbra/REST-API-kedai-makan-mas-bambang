package authentication

import "github.com/maulIbra/clean-architecture-go/api-master/models"

type IAuthenticationUsecase interface {
	AddUserProfile(profile *models.Profile) error
	ReadUserByEmail(username string) (*models.User, error)
	ReadUser() (*[]models.Profile, error)
	DeleteUser(userID string) error
}
