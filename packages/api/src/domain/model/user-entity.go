package model

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Name     UserName  `json:"name"`
	Password string    `json:"password"`
}
