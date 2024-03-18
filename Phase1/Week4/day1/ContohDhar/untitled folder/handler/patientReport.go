package handler

import (
	"fmt"
	"livecode-3-set-2-StevenDharmawan/config"
	"time"
)

func PatientReport() {
	query := "SELECT nama, descript FROM Patients"
	rows, err := config.DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var id int
		var name, medicalHistory string
		var dateOfBirth time.Time
		err = rows.Scan(&id, &name, &dateOfBirth, &medicalHistory)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("ID : ", id)
		fmt.Println("Name : ", name)
		fmt.Println("Date Of Birth : ", dateOfBirth)
		fmt.Println("Medical History : ", medicalHistory)
	}
}
