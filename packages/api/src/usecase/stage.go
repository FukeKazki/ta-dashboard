package usecase

import "github.com/FukeKazki/ta-dashboard/src/domain/model"

type StageUsecase interface {
	FindAll() ([]*model.Stage, error)
}

type stageUsecase struct {
}

func NewStageUsecase() StageUsecase {
	return &stageUsecase{}
}

func (u *stageUsecase) FindAll() ([]*model.Stage, error) {
	var stages = []*model.Stage{
		&model.Stage{
			ID:   1,
			Name: "1",
			TimeRecord: model.TimeRecord{
				FirstTime:  "00:00:00",
				SecondTime: "00:00:00",
				ThirdTime:  "00:00:00",
				TotalTime:  "00:00:00",
			},
		},
	}

	return stages, nil
}
