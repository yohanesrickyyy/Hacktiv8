package handler

import (
	"fmt"
	"livecode-3-set-2-StevenDharmawan/config"
)

func AppointmentStats() {
	query := "SELECT IFNULL(status, 'scheduled'), COUNT(*) FROM Treatments GROUP BY Status"
	rows, err := config.DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var status string
		var count int
		err = rows.Scan(&status, &count)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Status : ", status)
		fmt.Println("count : ", count)
	}
}
