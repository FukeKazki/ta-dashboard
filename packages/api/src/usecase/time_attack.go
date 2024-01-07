package usecase

import (
	"github.com/FukeKazki/ta-dashboard/src/domain/model"
	"github.com/FukeKazki/ta-dashboard/src/domain/repository"
)

type TimeAttackUsecase interface {
	FindUserTARecord(userName string) ([]*model.TimeAttack, error)
	CreateTARecord(timeAttack *model.TimeAttack, userName string) (*model.TimeAttack, error)
	UpdateTARecord(recordId int, timeAttack *model.TimeAttack) error
}

type timeAttackUsecase struct {
	timeAttackRepository repository.TimeAttackRepository
}

func NewTimeAttackUsecase(timeAttackRepo repository.TimeAttackRepository) TimeAttackUsecase {
	return &timeAttackUsecase{timeAttackRepository: timeAttackRepo}
}

func (u *timeAttackUsecase) FindUserTARecord(userName string) ([]*model.TimeAttack, error) {
	timeAttacks, err := u.timeAttackRepository.FindUserTARecord(userName)
	if err != nil {
		return nil, err
	}

	return timeAttacks, nil
}

func (u *timeAttackUsecase) CreateTARecord(timeAttack *model.TimeAttack, userName string) (*model.TimeAttack, error) {
	timeAttack, err := u.timeAttackRepository.CreateTARecord(timeAttack, userName)
	if err != nil {
		return nil, err
	}

	return timeAttack, nil
}

func (u *timeAttackUsecase) UpdateTARecord(recordId int, timeAttack *model.TimeAttack) error {
	err := u.timeAttackRepository.UpdateTARecord(recordId, timeAttack)
	if err != nil {
		return err
	}

	return nil
}
