package migration

import (
	"backend-smart-garden/config"
	"backend-smart-garden/entity"
	"fmt"
	"time"
)

func SeedGarden() {

	layoutTime := "2006-01-02 15:04:05"

	inputTime := "2024-04-16 19:38:40"

	loc, _ := time.LoadLocation("Asia/Jakarta")

	t, err := time.ParseInLocation(layoutTime, inputTime, loc)

	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	var garden = []entity.Garden{
		{
			NamaNode:    "Node01",
			Kelembapan:  51.6,
			TanggalNode: t,
		},
	}

	config.DB.Save(garden)
}
