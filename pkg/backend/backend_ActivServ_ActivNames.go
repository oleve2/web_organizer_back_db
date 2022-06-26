package backendServ

import (
	"database/sql"
	"fmt"
	"log"

	models "webapp3/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

// ---------------------------------------
// activities - names
func (s *Service) GetActivitiesList() ([]*models.ActivityDTO, error) {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("select id, name, date_start, date_end, norm_id from activ_names")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	activities := make([]*models.ActivityDTO, 0)

	for rows.Next() {
		tmp := &models.ActivityDTO{}
		err := rows.Scan(&tmp.Id, &tmp.Name, &tmp.DateStart, &tmp.DateEnd, &tmp.NormId)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		activities = append(activities, tmp)
	}
	//
	return activities, nil
}

//
func (s *Service) InsertNewActiv(newActiv *models.ActivityDTO) error {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("insert into activ_names (name, date_start, date_end, norm_id) values (?, ?, ?, ?)",
		newActiv.Name, newActiv.DateStart, newActiv.DateEnd, newActiv.NormId,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//
func (s *Service) UpdateNewActivById(updActiv *models.ActivityDTO) error {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()
	_, err = db.Exec("update activ_names set name=?, date_start=?, date_end=?, norm_id=? where id=?",
		updActiv.Name, updActiv.DateStart, updActiv.DateEnd, updActiv.NormId, updActiv.Id,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//
func (s *Service) DeleteNewActivById(delId string) error {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()
	_, err = db.Exec("delete from activ_names where id=?", delId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
