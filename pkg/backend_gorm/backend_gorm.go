package backendGormServ

import (
	"webapp3/pkg/models"

	"gorm.io/gorm"
)

type Service struct {
	sqliteDB *gorm.DB
}

func NewService(sqliteDB *gorm.DB) *Service {
	return &Service{
		sqliteDB: sqliteDB,
	}
}

/*
[Информация]
Альтернативный способ работы с бекендом - в дополнение к классическому на проекте

https://gorm.io/docs/models.html#gorm-Model
https://gorm.io/docs/transactions.html#Control-the-transaction-manually
https://stackoverflow.com/questions/23669720/meaning-of-interface-dot-dot-dot-interface

gorm youtube https://www.youtube.com/watch?v=9koLNdEcSR0

Вопрос - как обрабатывать ошибку при ответе gorm?
Ответ - присваивать в result
*/

func (s *Service) Init() {
	s.sqliteDB.AutoMigrate(
		&models.Tags{},
		&models.CalendarData{},
	)
}
