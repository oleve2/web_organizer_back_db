package backendServ

import (
	"database/sql"
	"fmt"
	"log"

	models "webapp3/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

// ---------------------------------------
// activities - normatives
func (s *Service) GetActivNorm() ([]*models.ActivityNormativeDTO, error) {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("select id, name, norm_period, norm_value, norm_measure from activ_normative")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	norms := make([]*models.ActivityNormativeDTO, 0)
	for rows.Next() {
		tmp := &models.ActivityNormativeDTO{}
		err := rows.Scan(&tmp.Id, &tmp.Name, &tmp.NormPeriod, &tmp.NormValue, &tmp.NormMeasure)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		norms = append(norms, tmp)
	}
	//
	return norms, nil
}

func (s *Service) ActivNormNew(newActivNorm *models.ActivityNormativeDTO) error {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("insert into activ_normative (name, norm_period, norm_value, norm_measure) values (?, ?, ?, ?)",
		newActivNorm.Name, newActivNorm.NormPeriod, newActivNorm.NormValue, newActivNorm.NormMeasure,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) ActivNormUpdate(updActivNorm *models.ActivityNormativeDTO) error {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()
	_, err = db.Exec("update activ_normative set name=?, norm_period=?, norm_value=?, norm_measure=? where id=?",
		updActivNorm.Name, updActivNorm.NormPeriod, updActivNorm.NormValue, updActivNorm.NormMeasure, updActivNorm.Id,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) ActivNormDelById(delId string) error {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()
	_, err = db.Exec("delete from activ_normative where id=?", delId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
