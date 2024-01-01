package model

import "database/sql"

// UserID    int    `json:"user_id"`
// CourseID  int    `json:"course_id"`
type Record struct {
	ID        int            `json:"id"`
	FirstLap  sql.NullString `json:"first_lap"`
	SecondLap sql.NullString `json:"second_lap"`
	ThirdLap  sql.NullString `json:"third_lap"`
	TotalLap  string         `json:"total_lap"`
}
