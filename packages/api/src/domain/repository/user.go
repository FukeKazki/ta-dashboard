package repository

import "github.com/FukeKazki/ta-dashboard/src/domain/model"

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
}
