package backendGormServ

import (
	"log"
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
Информация

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
		//&models.RandomTable{},
	)
}

// TAGS --------------------------------------------------------------------
func (s *Service) TagsAll() ([]*models.Tags, error) {
	var tagsArr []*models.Tags
	result := s.sqliteDB.Find(&models.Tags{}).Find(&tagsArr)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return tagsArr, nil
}

func (s *Service) TagsInsertOne(tagNew *models.Tags) error {
	result := s.sqliteDB.Create(&models.Tags{
		Name:  tagNew.Name,
		Color: tagNew.Color,
	})
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (s *Service) TagsUpdateOne(tagUpd *models.Tags) error {
	result := s.sqliteDB.Where(&models.Tags{ID: tagUpd.ID}).Updates(models.Tags{
		Name:  tagUpd.Name,
		Color: tagUpd.Color,
	})
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (s *Service) TagsDeleteOne(id int) error {
	result := s.sqliteDB.Delete(&models.Tags{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

// ??? --------------------------------------------------------------------
