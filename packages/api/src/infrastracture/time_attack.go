package infrastracture

import (
	"database/sql"

	"github.com/FukeKazki/ta-dashboard/src/domain/model"
)

type TimeAttackRepository struct {
	Connection *sql.DB
}

func NewTimeAttackRepository(conn *sql.DB) *TimeAttackRepository {
	return &TimeAttackRepository{Connection: conn}
}

func (r *TimeAttackRepository) FindUserTARecord(userId int) ([]*model.TimeAttack, error) {
	timeAttacks := make([]*model.TimeAttack, 0)
	rows, err := r.Connection.Query("select course.name as 'couse_name', first_lap, second_lap, third_lap, total_lap, thumbnail from record left outer join course on record.course_id=course.id where user_id=1;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var timeAttack model.TimeAttack
		err := rows.Scan(&timeAttack.Course.Name, &timeAttack.Record.FirstLap, &timeAttack.Record.SecondLap, &timeAttack.Record.ThirdLap, &timeAttack.Record.TotalLap, &timeAttack.Course.Thumbnail)
		if err != nil {
			return nil, err
		}
		timeAttacks = append(timeAttacks, &timeAttack)
	}

	return timeAttacks, nil
}
