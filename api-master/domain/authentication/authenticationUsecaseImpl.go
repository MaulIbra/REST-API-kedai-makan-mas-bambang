package authentication

import "github.com/maulIbra/clean-architecture-go/api-master/models"

type AuthenticationUsecase struct {
	AuthRepo IAuthenticationRepo
}

func NewAuthenticationUsecase(repo IAuthenticationRepo) IAuthenticationUsecase{
	return &AuthenticationUsecase{
		AuthRepo: repo,
	}
}

func (a AuthenticationUsecase) AddUserProfile(profile *models.Profile) error {
	error := a.AuthRepo.AddUserProfile(profile)
	if error != nil {
		return error
	}
	return nil
}

func (a AuthenticationUsecase) ReadUserByEmail(username string) (*models.User, error) {
	user,err := a.AuthRepo.ReadUserByEmail(username)
	if err!= nil {
		return nil, err
	}
	return user,nil
}

func (a AuthenticationUsecase) ReadUser() (*[]models.Profile, error) {
	listProfile,err := a.AuthRepo.ReadUser()
	if err != nil {
		return nil, err
	}
	return listProfile,nil
}

func (a AuthenticationUsecase) DeleteUser(userID string) error {
	err := a.AuthRepo.DeleteUser(userID)
	if err != nil {
		return err
	}
	return nil
}