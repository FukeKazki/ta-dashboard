package usecase

import (
	"github.com/FukeKazki/ta-dashboard/src/domain/model"
	"github.com/FukeKazki/ta-dashboard/src/domain/repository"
)

type UserUsecase interface {
	CreateUser(username string, password string) (*model.User, error)
	FindByName(name string) (*model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
	taRepo   repository.TimeAttackRepository
}

func NewUserUsecase(userRepo repository.UserRepository, taRepo repository.TimeAttackRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo, taRepo: taRepo}
}

func (u *userUsecase) CreateUser(username string, password string) (*model.User, error) {
	userName, err := model.NewUserName(username)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.Create(&model.User{
		Name:     *userName,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	// User作成時に全てのコースのタイムアタックを作成する
	err = u.taRepo.CreateAllTARecord(username)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) FindByName(name string) (*model.User, error) {
	user, err := u.userRepo.FindByName(name)
	if err != nil {
		return nil, err
	}

	return user, nil
}
