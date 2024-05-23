package repository

import (
	"backend-smart-garden-golang/config"
	"backend-smart-garden-golang/entity"
)

func SaveDataSoilMoistureSensor(sensor entity.SoilMoitureSensor) (entity.SoilMoitureSensor, error) {
	err := config.DB.Create(&sensor).Error
	if err != nil {
		return entity.SoilMoitureSensor{}, err
	}
	return sensor, nil
}
