CREATE DATABASE HealthTrack;

CREATE TABLE Patients (
	PatientId INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    DateOfBirth DATE,
    MedicalHistory TEXT
);

CREATE TABLE Doctors (
	DoctorId INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Speciality VARCHAR(255) NOT NULL,
    Availability varchar(255)
);

CREATE TABLE Appointments (
	AppointmentId INT AUTO_INCREMENT PRIMARY KEY,
    PatientId INT NOT NULL,
    DoctorId INT NOT NULL,
    `Date` DATE NOT NULL,
    `Time` TIME NOT NULL,
    FOREIGN KEY (PatientId) REFERENCES Patients(PatientId),
    FOREIGN KEY (DoctorId) REFERENCES Doctors(DoctorId)
);

CREATE TABLE Treatments (
	TreatmentId INT AUTO_INCREMENT PRIMARY KEY,
    PatientId INT NOT NULL,
    DoctorId INT NOT NULL,
    Diagnosis TEXT NOT NULL,
    Medication TEXT,
    Status VARCHAR(50),
    FOREIGN KEY (PatientId) REFERENCES Patients(PatientId),
    FOREIGN KEY (DoctorId) REFERENCES Doctors(DoctorId)
);

-- PatientId dapat dihilangkan karna pada tabel Treatments telah terapat PatientId
CREATE TABLE Billing (
	BillId INT PRIMARY KEY AUTO_INCREMENT,
    TreatmentId INT NOT NULL,
    Amount DECIMAL(10,2) NOT NULL,
    FOREIGN KEY (TreatmentId) REFERENCES Treatments(TreatmentId)
);

INSERT INTO Patients (Name, DateOfBirth, MedicalHistory) VALUES ('John Doe', '1990-05-15', 'No known allergies. Had chickenpox in 1995.');
INSERT INTO Patients (Name, DateOfBirth, MedicalHistory) VALUES ('Jane Smith', '1985-08-12', 'Allergic to penicillin. Broke left arm in 2005.');
INSERT INTO Patients (Name, DateOfBirth, MedicalHistory) VALUES ('Alice Johnson', '1999-12-03', 'Has asthma. No other known conditions.');

INSERT INTO Doctors (Name, Specialty, Availability) VALUES ('Dr. Peter Parker', 'Cardiologist', 'Monday-Wednesday');
INSERT INTO Doctors (Name, Specialty, Availability) VALUES ('Dr. Bruce Wayne', 'Orthopedic', 'Thursday-Saturday');
INSERT INTO Doctors (Name, Specialty, Availability) VALUES ('Dr. Clark Kent', 'General Practitioner', 'Monday-Friday');

INSERT INTO Appointments (PatientId, DoctorId, `Date`, `Time`) VALUES (1, 1, '2023-10-10', '10:00:00');
INSERT INTO Appointments (PatientId, DoctorId, `Date`, `Time`) VALUES (2, 2, '2023-10-11', '11:00:00');
INSERT INTO Appointments (PatientId, DoctorId, `Date`, `Time`) VALUES (3, 3, '2023-10-12', '14:00:00');
INSERT INTO Appointments (PatientId, DoctorId, `Date`, `Time`) VALUES (1, 1, '2024-10-10', '10:00:00');
INSERT INTO Appointments (PatientId, DoctorId, `Date`, `Time`) VALUES (2, 2, '2024-10-11', '11:00:00');
INSERT INTO Appointments (PatientId, DoctorId, `Date`, `Time`) VALUES (3, 3, '2024-10-12', '14:00:00');


INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication, `Status`) VALUES (1, 1, 'Minor chest pain due to stress.', 'Prescribed Ibuprofen.', 'Completed');
INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication, `Status`) VALUES (2, 2, 'Mild joint pain.', 'Prescribed Paracetamol.', 'Completed');
INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication, `Status`) VALUES (3, 3, 'Flu', 'Prescribed general flu medication.', 'Completed');
INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication, `Status`) VALUES (1, 1, 'Minor chest pain due to stress.', 'Prescribed Ibuprofen.', 'Missed');
INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication, `Status`) VALUES (2, 2, 'Mild joint pain.', 'Prescribed Paracetamol.', 'Missed');
INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication, `Status`) VALUES (3, 3, 'Flu', 'Prescribed general flu medication.', 'Missed');
INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication, `Status`) VALUES (3, 3, 'Flu', 'Prescribed general flu medication.', 'Missed');


INSERT INTO Billing (TreatmentId, Amount) VALUES (1, 150.00);
INSERT INTO Billing (TreatmentId, Amount) VALUES (2, 100.50);
INSERT INTO Billing (TreatmentId, Amount) VALUES (3, 80.25);
INSERT INTO Billing (TreatmentId, Amount) VALUES (4, 150.00);
INSERT INTO Billing (TreatmentId, Amount) VALUES (5, 100.50);
INSERT INTO Billing (TreatmentId, Amount) VALUES (6, 80.25);
INSERT INTO Billing (TreatmentId, Amount) VALUES (7, 80.25);