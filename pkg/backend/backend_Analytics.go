package backendServ

import (
	"database/sql"
	"fmt"
	"strconv"
	"webapp3/pkg/models"

	//models "webapp3/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

// -------------------------------------------------
// Activities - list of names
func (s *Service) GetAllActivNames() ([]*models.AnActivNameList, error) {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(`
	select T1.activ_name_id, T2.name 
	from activ_log as T1
	inner join activ_names as T2 on T1.activ_name_id = T2.id
	group by T1.activ_name_id, T2.name
	order by activ_name_id
	`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	activNameIdList := make([]*models.AnActivNameList, 0)
	for rows.Next() {
		tmp := &models.AnActivNameList{}
		err := rows.Scan(&tmp.Id, &tmp.Name)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		activNameIdList = append(activNameIdList, tmp)
	}
	//
	return activNameIdList, nil
}

// -------------------------------------------------
// Analytics - Activ3 - separate graphs for each activities
func (s *Service) AnalyticsActiv3(dateFrom string, dateTo string) ([]*models.AnNameLabelsValues, error) {
	// получить список активностей
	activNameList, err := s.GetAllActivNames()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// итоговый список
	listNLV := make([]*models.AnNameLabelsValues, 0)

	//
	for _, v := range activNameList {
		tmp := &models.AnNameLabelsValues{}
		tmp.ChartNameId = v.Id
		tmp.ChartName = v.Name

		labs, vals, err := s.GetDataByActivity(v.Id, dateFrom, dateTo, "activ3")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		tmp.Labels = labs
		tmp.Data = vals
		//
		listNLV = append(listNLV, tmp)
	}

	return listNLV, nil
}

// Support Activ3/Activ4  - данные для единичных графиков
func (s *Service) GetDataByActivity(actNameId int, dateFrom string, dateTo string, flgType string) ([]string, []int, error) {
	actNameIdStr := strconv.Itoa(actNameId)
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer db.Close()
	/*
		select activ_date, sum(activ_value)
		from activ_log
		where activ_name_id=%s
		and activ_date between '%s' and '%s'
		group by activ_date order by activ_date
	*/

	/*
		select T1.date, sum(coalesce(T2.activ_value, 0))
		from dates as T1
		left join (
			select activ_date, activ_value
			from activ_log
			where activ_name_id=%s
		) as T2 on T1.date=T2.activ_date
		where 1=1
		and T1.date between '%s' and '%s'
		group by date
		order by date asc
	*/

	// sql в зависимости от флага - для какого отчета гототвим данные
	var sql string
	if flgType == "activ3" {
		sql = fmt.Sprintf(`
select activ_date, sum(activ_value)
from activ_log
where activ_name_id=%s
and activ_date between '%s' and '%s'
group by activ_date order by activ_date
		`, actNameIdStr, dateFrom, dateTo)
	}
	if flgType == "activ4" {
		sql = fmt.Sprintf(`
select T1.date, sum(coalesce(T2.activ_value, 0))
from dates as T1
left join (
	select activ_date, activ_value
	from activ_log
	where activ_name_id=%s
) as T2 on T1.date=T2.activ_date
where 1=1
and T1.date between '%s' and '%s'
group by date
order by date asc
		`, actNameIdStr, dateFrom, dateTo)
	}

	rows, err := db.Query(sql)
	defer rows.Close()

	listLabels := make([]string, 0)
	listData := make([]int, 0)

	for rows.Next() {
		tmp := &models.AnDateVal{}
		err := rows.Scan(&tmp.ActivDate, &tmp.ActivValue)
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
		listLabels = append(listLabels, tmp.ActivDate)
		listData = append(listData, tmp.ActivValue)
	}
	//
	return listLabels, listData, nil
}

// -------------------------------------------------
// Analytics - Activ4 - combined graph with activities
func (s *Service) Activ4CommonGraphByDates(dateFrom string, dateTo string) (*models.ActivityDataReport, error) {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	// 01) получить общий список дат (Labels)
	sqlDates := fmt.Sprintf(`
select distinct date 
from dates 
where date between '%s' and '%s' 
order by date asc
	`, dateFrom, dateTo)

	rows, err := db.Query(sqlDates)
	defer rows.Close()

	ListOfDates := make([]string, 0)
	for rows.Next() {
		var tmp string
		err := rows.Scan(&tmp)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		ListOfDates = append(ListOfDates, tmp)
	}
	//fmt.Printf("ListOfDates = %+v \n", ListOfDates)

	// 02)	получить список активностей
	activNameList, err := s.GetAllActivNames()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 03) собрать инфу по активностям
	activityDataList := make([]*models.ActivityData, 0) // итоговый список

	for _, v := range activNameList {
		tmp := &models.ActivityData{}
		tmp.Label = v.Name // заполняем Label

		_, vals, err := s.GetDataByActivity(v.Id, dateFrom, dateTo, "activ4")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		tmp.Data = vals // заполняем Data
		//
		activityDataList = append(activityDataList, tmp) // добавляем в список
	}

	// 04) сборка финального ответа
	activ4Report := &models.ActivityDataReport{}
	activ4Report.Labels = ListOfDates
	activ4Report.Datasets = activityDataList

	//
	return activ4Report, nil
}
