package repository

import "github.com/FukeKazki/ta-dashboard/src/domain/model"

type TimeAttackRepository interface {
	FindUserTARecord(userId int) ([]*model.TimeAttack, error)
}
