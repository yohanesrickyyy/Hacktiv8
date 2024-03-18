package main

import (
	"fmt"
	"livecode-3-set-2-StevenDharmawan/config"
	"livecode-3-set-2-StevenDharmawan/handler"
)

func main() {
	config.ConnectDB()
	defer config.DB.Close()
	var inputUser int
	for {
		fmt.Println("--- Patient Management Interface ---")
		fmt.Println("1. Register a patient\n2. Add treatment\n3. Patient Reports\n4. Doctor Activity\n5. Billing Overview\n6. Appointment Stats\n7. Exit")
		fmt.Print("Masukkan pilihanmu: ")
		fmt.Scanln(&inputUser)
		switch inputUser {
		case 1:
			var name, dateOfBirth, medicalHistory string
			fmt.Print("Enter patient name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter Date of Birth (YYYY-MM-DD): ")
			fmt.Scanln(&dateOfBirth)
			fmt.Print("Enter Medical History: ")
			fmt.Scanln(&medicalHistory)
			handler.RegisterPatient(name, dateOfBirth, medicalHistory)
		case 2:
			var patientId, doctorId int
			var diagnosis, medication string
			fmt.Print("Enter patient id: ")
			fmt.Scanln(&patientId)
			fmt.Print("Enter doctor id: ")
			fmt.Scanln(&doctorId)
			fmt.Print("Enter diagnosis: ")
			fmt.Scanln(&diagnosis)
			fmt.Print("Enter medication: ")
			fmt.Scanln(&medication)
			handler.AddTreatment(patientId, doctorId, diagnosis, medication)
		case 3:
			handler.PatientReport()
		case 4:
			handler.DoctorActivity()
		case 5:
		case 6:
			handler.AppointmentStats()
		case 7:
			return
		}
	}
}
