package utils

import (
	"fmt"
	"time"
)

func StringDatetimeToTime(datetime string) time.Time {
	layoutTime := "2006-01-02 15:04:05"

	t, err := time.Parse(layoutTime, datetime)

	if err != nil {
		fmt.Println("Error parsing time:", err)
	}

	return t
}
