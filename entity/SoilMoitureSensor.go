package entity

import (
	"gorm.io/gorm"
	"time"
)

type SoilMoitureSensor struct {
	gorm.Model
	ID                uint      `gorm:"primaryKey;autoIncrement"`
	NilaiSensorAnalog uint      `gorm:"not null" json:"nilai_sensor_analog"`
	LabelSensorAnalog string    `gorm:"size:100; not null" json:"label_sensor_analog"`
	TanggalSensor     time.Time `gorm:"not null" json:"tanggal_sensor"`
}
