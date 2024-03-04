package backendServ

import (
	"database/sql"
	"fmt"
	"log"

	models "webapp3/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

// А что если тут воспользоваться goorm ? Напрягает написание большого кол-ва однотипного кода
// https://tutorialedge.net/golang/golang-orm-tutorial/

// заглушка :)
func (s *Service) A() (int, error) {
	return 1, nil
}

// ---------------------------------------
// activities - logs
func (s *Service) ActivLogsAll() ([]*models.ActivityLogDTO, error) {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
	select 
	T1.id, T1.activ_name_id, T1.activ_norm_id, T1.activ_date, T1.activ_value,
	T2.name
	from activ_log as T1
	inner join activ_names as T2 on T1.activ_name_id = T2.id
	`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	activLogs := make([]*models.ActivityLogDTO, 0)
	for rows.Next() {
		tmp := &models.ActivityLogDTO{}
		err := rows.Scan(&tmp.Id, &tmp.ActivNameId, &tmp.ActivNormId, &tmp.ActivDate, &tmp.ActivValue,
			&tmp.ActivName,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		activLogs = append(activLogs, tmp)
	}
	//
	return activLogs, nil
}

func (s *Service) ActivLogsNew(newActivLog *models.ActivityLogDTO) error {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("insert into activ_log (activ_name_id, activ_norm_id, activ_date, activ_value) values (?, ?, ?, ?)",
		newActivLog.ActivNameId, newActivLog.ActivNormId, newActivLog.ActivDate, newActivLog.ActivValue,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) ActivLogsUpdate(updActivLog *models.ActivityLogDTO) error {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("update activ_log set activ_name_id=?, activ_norm_id=?, activ_date=?, activ_value=? where id=?",
		updActivLog.ActivNameId, updActivLog.ActivNormId, updActivLog.ActivDate, updActivLog.ActivValue, updActivLog.Id,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) ActivLogsDelById(delId string) error {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("delete from activ_log where id=?", delId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
