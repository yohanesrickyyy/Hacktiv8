package handler

import (
	"fmt"
	"livecode-3-set-2-StevenDharmawan/config"
)

func DoctorActivity() {
	query := "SELECT COUNT(Doctors.DoctorId) FROM Appointments JOIN Doctors ON Appointments.DoctorId = Doctors.DoctorId GROUP BY Doctors.Name ORDER BY Doctors.DoctorId"
	rows, err := config.DB.Query(query)
	var appointment []int

	for rows.Next() {
		var count int
		err = rows.Scan(&count)
		if err != nil {
			panic(err.Error())
		}
		appointment = append(appointment, count)
	}

	query = "SELECT Doctors.Name, COUNT(Doctors.DoctorId) FROM Treatments JOIN Doctors ON Treatments.DoctorId = Doctors.DoctorId GROUP BY Doctors.Name ORDER BY Doctors.DoctorId"

	rows, err = config.DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
	counter := 0
	for rows.Next() {
		var name string
		var treatment int
		err = rows.Scan(&name, &treatment)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Doctor : ", name)
		fmt.Println("Treatment : ", treatment)

		fmt.Println("Appointment : ", appointment[counter])
		counter++
	}
}
