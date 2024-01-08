package model

import "database/sql"

type Record struct {
	ID        int            `json:"id"`
	FirstLap  sql.NullString `json:"first_lap"`
	SecondLap sql.NullString `json:"second_lap"`
	ThirdLap  sql.NullString `json:"third_lap"`
	TotalLap  string         `json:"total_lap"`
}

type Course struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}

type TimeAttack struct {
	Record Record `json:"record"`
	Course Course `json:"course"`
}
