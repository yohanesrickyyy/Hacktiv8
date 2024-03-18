package handler

import (
	"fmt"
	"livecode-3-set-2-StevenDharmawan/config"
)

func AddTreatment(patientId, doctorId int, diagnosis, medication string) error {
	query := "INSERT INTO Treatments(PatientID, DoctorID, Diagnosis, Medication) VALUES (?,?,?,?)"
	_, err := config.DB.Exec(query, patientId, doctorId, diagnosis, medication)
	if err != nil {
		return err
	}
	fmt.Println("Treatment registered successfully!")
	return nil
}
