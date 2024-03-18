package handler

import (
	"fmt"
	"livecode-3-set-2-StevenDharmawan/config"
)

func RegisterPatient(name string, dateOfBirth string, medicalHistory string) error {
	query := "INSERT INTO Patients(Name, DateOfBirth, MedicalHistory) VALUES (?,?,?)"
	_, err := config.DB.Exec(query, name, dateOfBirth, medicalHistory)
	if err != nil {
		return err
	}
	fmt.Println("Patient registered successfully!")
	return nil
}
