package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

// Database initialization and connection
func initializeDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "localhost:root@tcp(127.0.0.1:3306)/healthtrack")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Main function
func main() {
	db, err := initializeDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer db.Close()

	// Command line interface
	fmt.Println("Welcome to the Clinic Management System")
	fmt.Println("Please select an option:")
	fmt.Println("1. Register Patient")
	fmt.Println("2. Schedule Appointment")
	fmt.Println("3. Record Treatment")
	fmt.Println("4. Calculate Bill")
	fmt.Println("5. Patient Reports")
	fmt.Println("6. Doctor Activity")
	fmt.Println("7. Billing Overview")
	fmt.Println("8. Appointment Stats")
	fmt.Println("9. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		registerPatient(db)
	case 2:
		scheduleAppointment(db)
	case 3:
		recordTreatment(db)
	case 4:
		calculateBill(db)
	case 5:
		patientReports(db)
	case 6:
		doctorActivity(db)
	case 7:
		billingOverview(db)
	case 8:
		appointmentStats(db)
	case 9:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}

// Function to register a new patient
func registerPatient(db *sql.DB) {
	var name, dob, history string
	fmt.Println("Patient Registration:")
	fmt.Print("Enter patient name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter date of birth (YYYY-MM-DD): ")
	fmt.Scanln(&dob)
	fmt.Print("Enter medical history: ")
	fmt.Scanln(&history)

	_, err := db.Exec("INSERT INTO Patients (Name, DateOfBirth, MedicalHistory) VALUES (?, ?, ?)", name, dob, history)
	if err != nil {
		log.Println("Error registering patient:", err)
		return
	}

	fmt.Println("Patient registered successfully!")
}

// Function to schedule appointment
func scheduleAppointment(db *sql.DB) {
	var patientID, doctorID int
	var date, time string
	fmt.Println("Appointment Scheduling:")
	fmt.Print("Enter patient ID: ")
	fmt.Scanln(&patientID)
	fmt.Print("Enter doctor ID: ")
	fmt.Scanln(&doctorID)
	fmt.Print("Enter date (YYYY-MM-DD): ")
	fmt.Scanln(&date)
	fmt.Print("Enter time: ")
	fmt.Scanln(&time)

	_, err := db.Exec("INSERT INTO Appointments (PatientID, DoctorID, Date, Time) VALUES (?, ?, ?, ?)", patientID, doctorID, date, time)
	if err != nil {
		log.Println("Error scheduling appointment:", err)
		return
	}

	fmt.Println("Appointment scheduled successfully!")
}

// Function to record treatment
func recordTreatment(db *sql.DB) {
	var patientID, doctorID int
	var diagnosis, medication string
	fmt.Println("Recording Treatment:")
	fmt.Print("Enter patient ID: ")
	fmt.Scanln(&patientID)
	fmt.Print("Enter doctor ID: ")
	fmt.Scanln(&doctorID)
	fmt.Print("Enter diagnosis: ")
	fmt.Scanln(&diagnosis)
	fmt.Print("Enter medication: ")
	fmt.Scanln(&medication)

	_, err := db.Exec("INSERT INTO Treatments (PatientID, DoctorID, Diagnosis, Medication) VALUES (?, ?, ?, ?)", patientID, doctorID, diagnosis, medication)
	if err != nil {
		log.Println("Error recording treatment:", err)
		return
	}

	fmt.Println("Treatment recorded successfully!")
}

func calculateBill(db *sql.DB) {
	var patientID int
	fmt.Println("Billing Integration:")
	fmt.Print("Enter patient ID: ")
	fmt.Scanln(&patientID)

	rows, err := db.Query("SELECT SUM(Amount) FROM Billing WHERE PatientID=?", patientID)
	if err != nil {
		log.Println("Error calculating bill:", err)
		return
	}
	defer rows.Close()

	var totalAmount float64
	for rows.Next() {
		err := rows.Scan(&totalAmount)
		if err != nil {
			log.Println("Error scanning rows:", err)
			return
		}
	}

	fmt.Printf("Total amount to be paid by patient ID %d: $%.2f\n", patientID, totalAmount)
}

func patientReports(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM Patients")
	if err != nil {
		log.Println("Error fetching patient details:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Patient Reports:")
	fmt.Println("ID\tName\tDate of Birth\tMedical History")
	for rows.Next() {
		var id int
		var name, dob, history string
		err := rows.Scan(&id, &name, &dob, &history)
		if err != nil {
			log.Println("Error scanning rows:", err)
			return
		}
		fmt.Printf("%d\t%s\t%s\t%s\n", id, name, dob, history)
	}
}

func doctorActivity(db *sql.DB) {
	rows, err := db.Query("SELECT Doctors.Name, COUNT(Appointments.ID)")
	if err != nil {
		log.Println("Error fetching doctor activity:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Doctor Activity:")
	fmt.Println("Name\tAppointments\tTreatments")
	for rows.Next() {
		var name string
		var appointmentsCount, treatmentsCount int
		err := rows.Scan(&name, &appointmentsCount, &treatmentsCount)
		if err != nil {
			log.Println("Error scanning rows:", err)
			return
		}
		fmt.Printf("%s\t%d\t\t%d\n", name, appointmentsCount, treatmentsCount)
	}
}

func billingOverview(db *sql.DB) {
	var totalRevenue, outstandingBills, dailyIncome float64

	row := db.QueryRow("SELECT SUM(Amount) FROM Billing")
	row.Scan(&totalRevenue)

	row = db.QueryRow("SELECT SUM(Amount) FROM Billing WHERE Paid = 0")
	row.Scan(&outstandingBills)

	row = db.QueryRow("SELECT SUM(Amount) FROM Billing WHERE Date = '2024-03-12'")
	row.Scan(&dailyIncome)

	fmt.Printf("Billing Overview:\nTotal Revenue: $%.2f\nOutstanding Bills: $%.2f\nDaily Income: $%.2f\n", totalRevenue, outstandingBills, dailyIncome)
}

func appointmentStats(db *sql.DB) {
	var scheduled, completed, missed int

	row := db.QueryRow("SELECT COUNT(*) FROM Appointments")
	row.Scan(&scheduled)

	row = db.QueryRow("SELECT COUNT(*) FROM Appointments WHERE Completed = 1")
	row.Scan(&completed)

	// Missed appointments
	row = db.QueryRow("SELECT COUNT(*) FROM Appointments WHERE Completed = 0")
	row.Scan(&missed)

	fmt.Printf("Appointment Statistics:\nScheduled: %d\nCompleted: %d\nMissed: %d\n", scheduled, completed, missed)
}