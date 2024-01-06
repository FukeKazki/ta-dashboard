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
	// insert public_profile
	publicProfile, err := r.Connection.Exec("INSERT INTO public_profile (NAME) VALUES (?)", user.Name)
	if err != nil {
		return nil, err
	}
	// get public_profile_id
	publicProfileID, err := publicProfile.LastInsertId()
	if err != nil {
		return nil, err
	}
	// binary uuid   id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
	uu := uuid.New()
	userID, err := uu.MarshalBinary()

	// insert user
	_, err = r.Connection.Exec("INSERT INTO user (ID, PUBLIC_PROFILE_ID, NAME, PASSWORD) VALUES (?, ?, ?, ?)", userID, publicProfileID, user.Name, user.Password)
	if err != nil {
		return nil, err
	}
	return &model.User{Id: uu, Name: user.Name, Password: user.Password}, nil

}
