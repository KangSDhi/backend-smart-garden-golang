package repository

import (
	"backend-smart-garden/config"
	"backend-smart-garden/entity"
)

func SaveDataGarden(garden entity.Garden) (entity.Garden, error) {
	err := config.DB.Create(&garden).Error
	if err != nil {
		return entity.Garden{}, err
	}
	return garden, err
}
