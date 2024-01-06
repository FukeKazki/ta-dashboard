package usecase

import (
	"github.com/FukeKazki/ta-dashboard/src/domain/model"
	"github.com/FukeKazki/ta-dashboard/src/domain/repository"
)

type UserUsecase interface {
	CreateUser(username string, password string) (*model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) CreateUser(username string, password string) (*model.User, error) {
	user, err := u.userRepo.Create(&model.User{Name: username, Password: password})
	// TODO: User作成時に
	if err != nil {
		return nil, err
	}

	return user, nil
}
