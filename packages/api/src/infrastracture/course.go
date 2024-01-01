package infrastracture

import (
	"database/sql"

	"github.com/FukeKazki/ta-dashboard/src/domain/model"
)

type CourseRepository struct {
	Connection *sql.DB
}

func NewCourseRepository(conn *sql.DB) *CourseRepository {
	return &CourseRepository{Connection: conn}
}

func (r *CourseRepository) FindAll() ([]*model.Course, error) {
	courses := make([]*model.Course, 0)
	rows, err := r.Connection.Query("SELECT ID, NAME, THUMBNAIL FROM course")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var course model.Course
		err := rows.Scan(&course.ID, &course.Name, &course.Thumbnail)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}

	return courses, nil
}
