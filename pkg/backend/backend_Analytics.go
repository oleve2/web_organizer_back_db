package backendServ

import (
	"database/sql"
	"fmt"
	"webapp3/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

// -------------------------------------------------
func (s *Service) PrepareCommonGraphs(dateFrom string, dateTo string) ([]*models.AnNameLabelsValues, error) {
	// получить список активностей
	activNameList, err := s.SUPP_GetAllActivNames()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// итоговый список
	listCommGraphs := make([]*models.AnNameLabelsValues, 0)

	//
	for _, v := range activNameList {
		tmp := &models.AnNameLabelsValues{}
		tmp.ChartNameId = v.Id
		tmp.ChartName = v.Name

		labs, vals, err := s.SUPP_GetDataByActivity(v.Id, dateFrom, dateTo, "common")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		// проверка на нулевую сумму
		sumVals := s.SUPP_sumIntList(vals)
		if sumVals > 0 {
			tmp.Labels = labs
			tmp.Data = vals
			//
			listCommGraphs = append(listCommGraphs, tmp)
		}
	}

	return listCommGraphs, nil
}

// -------------------------------------------------
func (s *Service) PrepareIndivGraphs(dateFrom string, dateTo string) (*models.ActivityDataReport, error) {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	// 01) получить общий список дат (Labels)
	ListOfDates, err := s.SUPP_getDates(dateFrom, dateTo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 02)	получить список активностей
	activNameList, err := s.SUPP_GetAllActivNames()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 03) собрать инфу по активностям
	activityDataList := make([]*models.ActivityData, 0) // итоговый список

	for _, v := range activNameList {
		tmp := &models.ActivityData{}
		tmp.Label = v.Name // заполняем Label

		_, vals, err := s.SUPP_GetDataByActivity(v.Id, dateFrom, dateTo, "individual")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		// проверка на нулевую сумму
		sumVals := s.SUPP_sumIntList(vals)
		if sumVals > 0 {
			tmp.Data = vals // заполняем Data
			//
			activityDataList = append(activityDataList, tmp) // добавляем в список
		}
	}

	// 04) сборка финального ответа
	activ4Report := &models.ActivityDataReport{}
	activ4Report.Labels = ListOfDates
	activ4Report.Datasets = activityDataList

	//
	return activ4Report, nil
}
