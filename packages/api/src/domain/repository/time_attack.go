package repository

import "github.com/FukeKazki/ta-dashboard/src/domain/model"

type TimeAttackRepository interface {
	FindUserTARecord(userName string) ([]*model.TimeAttack, error)
	UpdateTARecord(recordId int, timeAttack *model.TimeAttack) error
	CreateTARecord(timeAttack *model.TimeAttack, publicProfileId int) (*model.TimeAttack, error)
}
