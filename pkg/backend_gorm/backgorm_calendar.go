package backendGormServ

import (
	"fmt"
	"log"
	"math"
	"sort"
	"time"
	"webapp3/pkg/models"
)

// I) Calendar - CRUD -------------------------------------
func (s *Service) CalendItemAll(yearmonth string) ([]*models.CalendarData, error) {
	var ciArr []*models.CalendarData

	// month stare/end and weekdays (смотреть сигнатуру)
	m1b, mc, m1a := s.Calend_Get3YMWithCurrentInMiddle(yearmonth)
	//fmt.Println(m1b, mc, m1a)

	result := s.sqliteDB.
		Where("date like ? or date like ? or date like ?",
			fmt.Sprintf("%%%s%%", m1b), fmt.Sprintf("%%%s%%", mc), fmt.Sprintf("%%%s%%", m1a)).
		Find(&models.CalendarData{}).
		Find(&ciArr)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return ciArr, nil
}

func (s *Service) CalendItemsAllNoFilter() ([]*models.CalendarData, error) {
	var ciArr []*models.CalendarData
	result := s.sqliteDB.
		Order("date asc").
		Find(&models.CalendarData{}).
		Find(&ciArr)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return ciArr, nil
}

func (s *Service) CalendItemInsertOne(calendItemNew *models.CalendarData) error {
	result := s.sqliteDB.Create(&models.CalendarData{
		Date:     calendItemNew.Date,
		Name:     calendItemNew.Name,
		TimeFrom: calendItemNew.TimeFrom,
		TimeTo:   calendItemNew.TimeTo,
		Status:   calendItemNew.Status,
	})
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (s *Service) CalendItemUpdateOne(calendItemUpd *models.CalendarData) error {
	result := s.sqliteDB.
		Where(&models.CalendarData{ID: calendItemUpd.ID}).
		Updates(models.CalendarData{
			Date:     calendItemUpd.Date,
			Name:     calendItemUpd.Name,
			TimeFrom: calendItemUpd.TimeFrom,
			TimeTo:   calendItemUpd.TimeTo,
			Status:   calendItemUpd.Status,
		})
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (s *Service) CalendItemDeleteOne(id int) error {
	result := s.sqliteDB.Delete(&models.CalendarData{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

// II) Calendar - Logic -------------------------------------

func (s *Service) Calend_GetGrid(yearmonth string) [][]models.CalendarItem {
	timeArr := s.Calend_F01_PrepareBaseArr(yearmonth) //"2024-02"
	timeArr2D := s.Calend_F02_Prepare2DArr(timeArr, yearmonth)

	return timeArr2D
}

func (s *Service) Calend_findcalendEventForDate(date string, dataArr []*models.CalendarData) []models.CalendarData {
	// получить массив
	res := make([]models.CalendarData, 0)
	for _, v := range dataArr {
		if v.Date == date {
			res = append(res, *v)
		}
	}
	// отсортировать по TimeFrom по возрастанию
	sort.Slice(res, func(i, j int) bool {
		return res[i].TimeFrom < res[j].TimeFrom
	})
	// вернуть массив
	return res
}

func (s *Service) Calend_makeMonthStartEndandWeekdays(yearmonth string) (time.Time, time.Time) {
	// первое и последнее числа месяца
	monthStart, _ := time.Parse("2006-01", yearmonth)
	monthEnd := monthStart.AddDate(0, 1, -1)
	return monthStart, monthEnd
}

func (s *Service) Calend_Get3YMWithCurrentInMiddle(yearmonth string) (string, string, string) {
	monthStart, _ := time.Parse("2006-01", yearmonth)
	month_1_Before := monthStart.AddDate(0, -1, 0)
	month_1_After := monthStart.AddDate(0, 1, 0)
	return month_1_Before.String()[0:7], monthStart.String()[0:7], month_1_After.String()[0:7]
}

func (s *Service) Calend_F01_PrepareBaseArr(yearmonth string) []string {
	//
	monthStart, monthEnd := s.Calend_makeMonthStartEndandWeekdays(yearmonth)
	monthStart_weekday := int(monthStart.Weekday())
	if monthStart_weekday == 0 {
		monthStart_weekday = 7
	}
	monthEnd_weekday := int(monthEnd.Weekday())
	//fmt.Println("monthStart=", monthStart, "weekday=", monthStart_weekday)
	//fmt.Println("monthEnd=", monthEnd, "weekday=", monthEnd_weekday)

	// tail
	tail := make([]time.Time, 0)
	for i := monthStart_weekday - 1; i >= 1; i-- {
		v := monthStart.AddDate(0, 0, -i)
		tail = append(tail, v)
	}

	// middle
	middle := make([]time.Time, 0)
	for i := monthStart.Day(); i <= monthEnd.Day(); i++ {
		v := monthStart.AddDate(0, 0, i-1) // начинаем прибавлять с нуля
		middle = append(middle, v)
	}

	// head
	head := make([]time.Time, 0)
	for i := monthEnd_weekday + 1; i <= 7; i++ {
		v := monthEnd.AddDate(0, 0, i-monthEnd_weekday)
		head = append(head, v)
	}

	// concatenate silces
	r := append(tail, middle...)
	r = append(r, head...)

	// convert time.Time to string
	rStr := make([]string, 0)
	for _, v := range r {
		v1 := v.String()[0:10]
		rStr = append(rStr, v1)
	}
	//
	return rStr
}

func (s *Service) Calend_F02_Prepare2DArr(timeArr []string, yearmonth string) [][]models.CalendarItem {
	// кол-во итераций
	timeArrLen := math.Ceil(float64(len(timeArr)) / 7.0)
	timeArrLen2 := int(timeArrLen)

	// получить список CalendarItem из БД
	dataCalendItems, err := s.CalendItemAll(yearmonth)
	if err != nil {
		log.Println("error retrieving CalendItemAll")
		return nil
	}

	// итог
	resArrWithCalendItems := make([][]models.CalendarItem, 0)
	for i := 0; i < timeArrLen2; i++ {
		oneWeekRow := make([]models.CalendarItem, 0)
		bottom := i * 7
		top := i*7 + 7
		// корректируем границы разбиения
		if top > len(timeArr) {
			top = len(timeArr)
		}
		//fmt.Println(i, bottom, top)
		subArr := timeArr[bottom:top]
		for _, v := range subArr {
			tmp := models.CalendarItem{
				ItemId: v,
				Data:   s.Calend_findcalendEventForDate(v, dataCalendItems),
			}
			oneWeekRow = append(oneWeekRow, tmp)
		}
		resArrWithCalendItems = append(resArrWithCalendItems, oneWeekRow)
	}

	return resArrWithCalendItems
}
