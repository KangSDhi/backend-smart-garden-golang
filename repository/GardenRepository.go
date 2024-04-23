package repository

import (
	"backend-smart-garden-golang/config"
	"backend-smart-garden-golang/entity"
)

func SaveDataGarden(garden entity.Garden) (entity.Garden, error) {
	err := config.DB.Create(&garden).Error
	if err != nil {
		return entity.Garden{}, err
	}
	return garden, err
}

func GetLastRecordGarden() (entity.Garden, error) {
	var garden entity.Garden
	err := config.DB.Last(&garden).Error
	if err != nil {
		return entity.Garden{}, err
	}
	return garden, nil
}
