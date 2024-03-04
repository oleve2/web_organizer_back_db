package backendServ

import (
	"database/sql"
	"fmt"
	"strconv"
	models "webapp3/pkg/models"
)

// -------------------------------------------------
// Support - получить список дат за период
func (s *Service) SUPP_getDates(dateFrom string, dateTo string) ([]string, error) {
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

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
	//
	return ListOfDates, nil
}

// -------------------------------------------------
// Support - получить список активностей
func (s *Service) SUPP_GetAllActivNames() ([]*models.AnActivNameList, error) {
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

// Support - данные для единичных графиков
/*
Описание параметра flgType:
	"common" - общий график (было "activ3")
	"individual" - индивидуальные графики (было "activ4")

	Детали отбора - возвращаем все что 	нашли в БД, отфильтровка нулей будет на уровень выше
*/
func (s *Service) SUPP_GetDataByActivity(actNameId int, dateFrom string, dateTo string, flgType string) ([]string, []int, error) {
	//
	actNameIdStr := strconv.Itoa(actNameId)
	db, err := sql.Open("sqlite3", s.sqlitePath)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer db.Close()

	// sql в зависимости от флага - для какого отчета гототвим данные
	var sql string

	// common
	if flgType == "common" {
		sql = fmt.Sprintf(`
select activ_date, sum(activ_value)
from activ_log
where activ_name_id=%s
and activ_date between '%s' and '%s'
group by activ_date order by activ_date
		`, actNameIdStr, dateFrom, dateTo)
	}

	// individual
	if flgType == "individual" {
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
// Support - проверка
func (s *Service) SUPP_sumIntList(intList []int) int {
	var sum int = 0
	for _, v := range intList {
		sum += v
	}
	return sum
}
