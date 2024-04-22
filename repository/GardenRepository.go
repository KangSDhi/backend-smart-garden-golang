package repository

import (
	"backend-smart-garden/config"
	"backend-smart-garden/entity"
	"log"
	"time"
)

func SaveDataGarden(garden entity.Garden) (entity.Garden, error) {
	startTime := time.Now()
	result, errDB := config.DB.Exec("INSERT INTO garden (nama_node, kelembapan, tanggal_node) VALUES (?, ?, ?)", garden.NamaNode, garden.Kelembapan, garden.TanggalNode)
	if errDB != nil {
		return entity.Garden{}, errDB
	}
	elapsedTime := time.Since(startTime)
	log.Printf("Query Insert Time : %s", elapsedTime)

	id, errID := result.LastInsertId()
	if errID != nil {
		return garden, errID
	}

	garden.ID = uint(id)

	return garden, nil
}

func GetLastRecordGarden() (entity.Garden, error) {
	var garden entity.Garden

	startTime := time.Now()
	row := config.DB.QueryRow("SELECT * FROM garden ORDER BY id DESC LIMIT 1")
	elapsedTime := time.Since(startTime)
	log.Printf("Query Read Time : %s", elapsedTime)
	if err := row.Scan(&garden.ID, &garden.NamaNode, &garden.Kelembapan, &garden.TanggalNode); err != nil {
		return entity.Garden{}, err
	}
	return garden, nil
}
