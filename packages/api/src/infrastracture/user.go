package infrastracture

import (
	"database/sql"

	"github.com/FukeKazki/ta-dashboard/src/domain/model"
	"github.com/google/uuid"
)

type UserRepository struct {
	Connection *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{Connection: conn}
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	// binary uuid   id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
	uu := uuid.New()
	userID, err := uu.MarshalBinary()

	// insert user
	_, err = r.Connection.Exec("INSERT INTO user (ID, NAME, PASSWORD) VALUES (?, ?, ?)", userID, user.Name, user.Password)
	if err != nil {
		return nil, err
	}
	return &model.User{Id: uu, Name: user.Name, Password: user.Password}, nil

}
