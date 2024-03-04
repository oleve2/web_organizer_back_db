package backendGormServ

import (
	"log"
	"webapp3/pkg/models"
)

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
	result := s.sqliteDB.
		Where(&models.Tags{ID: tagUpd.ID}).
		Updates(models.Tags{
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
