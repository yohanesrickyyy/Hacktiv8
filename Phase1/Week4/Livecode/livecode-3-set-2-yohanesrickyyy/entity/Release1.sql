-- Membuat database
CREATE DATABASE IF NOT EXISTS HealthTrack

-- Membuat Table 
CREATE TABLE Patients (
  PatientsID INT(11) PRIMARY KEY AUTO_INCREMENT,
  Name VARCHAR(255) NOT NULL,
  DateOfBirth VARCHAR(255) NOT NULL,
  MedicalHistory VARCHAR(255) NOT NULL
);

CREATE TABLE Doctors (
  DoctorsID INT(11) PRIMARY KEY AUTO_INCREMENT,
  Name VARCHAR(255) NOT NULL,
  Speciality VARCHAR(255) NOT NULL,
  Availability VARCHAR(255) NOT NULL
);

CREATE TABLE Appointments (
  AppointmentID INT(11) PRIMARY KEY AUTO_INCREMENT,
  PatientID INT,
  DoctorID INT,
  AppointmentDate DATE NOT NULL,
  AppointmentTime TIME NOT NULL,
  FOREIGN KEY (PatientID) REFERENCES Patients(PatientsID),
  FOREIGN KEY (DoctorID) REFERENCES Doctors(DoctorsID)
);

CREATE TABLE Treatement (
  TreatmentID INT(11) PRIMARY KEY AUTO_INCREMENT,
  PatientID INT NOT NULL,
  DoctorID INT NOT NULL,
  Diagnosis VARCHAR(255) NOT NULL,
  Medication VARCHAR(255) NOT NULL,
  FOREIGN KEY (PatientID) REFERENCES Patients(PatientsID),
  FOREIGN KEY (DoctorID) REFERENCES Doctors(DoctorsID)
);

CREATE TABLE Billing(
  BillID INT(11) PRIMARY KEY AUTO_INCREMENT,
  TreatmentID INT NOT NULL,
  Amount INT NOT NULL,
  FOREIGN KEY (TreatmentID) REFERENCES Treatement(TreatmentID)
);

-- Isi Table
INSERT INTO Patients (Name, DateOfBirth, MedicalHistory) 
VALUES 
('John Doe', '1990-05-15', 'No known allergies. Had chickenpox in 1995.'),
('Jane Smith', '1985-08-12', 'Allergic to penicillin. Broke left arm in 2005.'),
('Alice Johnson', '1999-12-03', 'Has asthma. No other known conditions.');

INSERT INTO Doctors (Name, Specialty, Availability) 
VALUES 
('Dr. Peter Parker', 'Cardiologist', 'Monday-Wednesday'),
('Dr. Bruce Wayne', 'Orthopedic', 'Thursday-Saturday'),
('Dr. Clark Kent', 'General Practitioner', 'Monday-Friday');

INSERT INTO Appointments (PatientId, DoctorId, Date, Time) 
VALUES 
(1, 1, '2023-10-10', '10:00:00'),
(2, 2, '2023-10-11', '11:00:00'),
(3, 3, '2023-10-12', '14:00:00');

INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication) 
VALUES 
(1, 1, 'Minor chest pain due to stress.', 'Prescribed Ibuprofen.'),
(2, 2, 'Mild joint pain.', 'Prescribed Paracetamol.'),
(3, 3, 'Flu', 'Prescribed general flu medication.');

INSERT INTO Billing (PatientId, TreatmentId, Amount) 
VALUES 
(1, 1, 150.00),
(2, 2, 100.50),
(3, 3, 80.25),