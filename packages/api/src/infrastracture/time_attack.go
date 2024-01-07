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

func (r *TimeAttackRepository) FindUserTARecord(userName string) ([]*model.TimeAttack, error) {
	timeAttacks := make([]*model.TimeAttack, 0)

	rows, err := r.Connection.Query("select course.name as 'couse_name', first_lap, second_lap, third_lap, total_lap, thumbnail from record left outer join course on record.course_id=course.id left outer join user on user.name=record.user_name where user_name=?", userName)
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

func (r *TimeAttackRepository) CreateTARecord(timeAttack *model.TimeAttack, userName string) (*model.TimeAttack, error) {
	_, err := r.Connection.Exec("insert into record (course_id, user_name, first_lap, second_lap, third_lap, total_lap) values (?, ?, ?, ?, ?, ?)", timeAttack.Course.ID, userName, timeAttack.Record.FirstLap, timeAttack.Record.SecondLap, timeAttack.Record.ThirdLap, timeAttack.Record.TotalLap)
	if err != nil {
		return nil, err
	}
	return timeAttack, nil
}

func (r *TimeAttackRepository) UpdateTARecord(recordId int, timeAttack *model.TimeAttack) error {
	_, err := r.Connection.Exec("update record set first_lap=?, second_lap=?, third_lap=?, total_lap=? where id=?", timeAttack.Record.FirstLap, timeAttack.Record.SecondLap, timeAttack.Record.ThirdLap, timeAttack.Record.TotalLap, recordId)
	if err != nil {
		return err
	}

	return nil
}

func (r *TimeAttackRepository) CreateAllTARecord(userName string) error {
	rows, err := r.Connection.Query("select id from course")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		courseId := 0
		err := rows.Scan(&courseId)
		if err != nil {
			return err
		}
		_, err = r.Connection.Exec("insert into record (course_id, user_name, total_lap) values (?, ?, ?)", courseId, userName, "00:00.00")
	}
	return nil
}
