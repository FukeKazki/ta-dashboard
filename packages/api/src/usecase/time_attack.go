package usecase

import (
	"github.com/FukeKazki/ta-dashboard/src/domain/model"
	"github.com/FukeKazki/ta-dashboard/src/domain/repository"
)

type TimeAttackUsecase interface {
	FindUserTARecord() ([]*model.TimeAttack, error)
}

type timeAttackUsecase struct {
	timeAttackRepository repository.TimeAttackRepository
}

func NewTimeAttackUsecase(timeAttackRepo repository.TimeAttackRepository) TimeAttackUsecase {
	return &timeAttackUsecase{timeAttackRepository: timeAttackRepo}
}

func (u *timeAttackUsecase) FindUserTARecord() ([]*model.TimeAttack, error) {
	timeAttacks, err := u.timeAttackRepository.FindUserTARecord(1)
	if err != nil {
		return nil, err
	}

	return timeAttacks, nil
}
