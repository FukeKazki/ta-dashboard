package repository

import "github.com/FukeKazki/ta-dashboard/src/domain/model"

// TODO:infrastractureのCourseRepositoryとの違いはなんだ。
type CourseRepository interface {
	FindAll() ([]*model.Course, error)
}
