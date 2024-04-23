package entity

import (
	"gorm.io/gorm"
	"time"
)

type Garden struct {
	gorm.Model
	ID          uint      `gorm:"primary_key"`
	NamaNode    string    `gorm:"size:255; not null;" json:"nama_node"`
	Kelembapan  float32   `gorm:"not null;" json:"kelembapan"`
	TanggalNode time.Time `gorm:"not null;" json:"tanggal_node"`
}
