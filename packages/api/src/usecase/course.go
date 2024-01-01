package usecase

import (
	"github.com/FukeKazki/ta-dashboard/src/domain/model"
	"github.com/FukeKazki/ta-dashboard/src/domain/repository"
)

type CourseUsecase interface {
	FindAll() ([]*model.Course, error)
}

type courseUsecase struct {
	courseRepository repository.CourseRepository
}

func NewCourseUsecase(courseRepo repository.CourseRepository) CourseUsecase {
	return &courseUsecase{courseRepository: courseRepo}
}

func (u *courseUsecase) FindAll() ([]*model.Course, error) {
	courses, err := u.courseRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return courses, nil
}
