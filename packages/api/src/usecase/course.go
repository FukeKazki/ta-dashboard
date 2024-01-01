package usecase

import "github.com/FukeKazki/ta-dashboard/src/domain/model"

type CourseUsecase interface {
	FindAll() ([]*model.Course, error)
}

type courseUsecase struct {
}

func NewCourseUsecase() CourseUsecase {
	return &courseUsecase{}
}

func (u *courseUsecase) FindAll() ([]*model.Course, error) {
	var courses = []*model.Course{
		&model.Course{
			ID:        1,
			Name:      "1",
			Thumbnail: "",
		},
	}

	return courses, nil
}
