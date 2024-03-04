package models

import "time"

type CalendarYearMonthDTO struct {
	YearMonth string `json:"year_month"`
}

type CalendarData struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Date      string    `json:"date"`
	Name      string    `json:"name"`
	TimeFrom  string    `json:"time_from"`
	TimeTo    string    `json:"time_to"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// форма для отпавки на фронтенд
type CalendarItem struct {
	ItemId string         `json:"item_id"`
	Data   []CalendarData `json:"data"`
}
