CREATE TABLE student (
  student_id SERIAL INTEGER  PRIMARY KEY,
  student_name VARCHAR(100) NOT NULL 
)

CREATE TABLE professor (
  professor_id SERIAL INTEGER PRIMARY KEY,
  professor_name VARCHAR(100) NOT NULL
)

CREATE TABLE course (
  course_id SERIAL INTEGER PRIMARY KEY,
  course_title VARCHAR(100) NOT NULL,
  max_capacity INT
)

CREATE TABLE Enrollment (
    enrollment_id SERIAL PRIMARY KEY,
    student_id INT REFERENCES Student(student_id),
    course_id INT REFERENCES Course(course_id),
    enrollment_date DATE
);

CREATE TABLE TeachingAssignment (
    assignment_id SERIAL PRIMARY KEY,
    professor_id INT REFERENCES Professor(professor_id),
    course_id INT REFERENCES Course(course_id),
    start_date DATE
);

-- Insert sample students
INSERT INTO Student (student_name) VALUES
('John Doe'),
('Jane Smith'),
('Alice Johnson');

-- Insert sample professors
INSERT INTO Professor (professor_name) VALUES
('Prof. Budi'),
('Prof. Indah'),
('Prof. Andreas'),
('Prof. John'),
('Prof. Tasya');

-- Insert sample courses
INSERT INTO Course (course_title, max_capacity) VALUES
('Introduction to Programming', 30),
('Database Management', 25),
('Golang', 40);

-- Insert sample enrollments
INSERT INTO Enrollment (student_id, course_id, enrollment_date) VALUES
(1, 1, CURRENT_DATE),
(2, 1, CURRENT_DATE),
(3, 3, CURRENT_DATE);

-- Insert sample teaching assignments
INSERT INTO TeachingAssignment (professor_id, course_id, start_date) VALUES
(1, 1, CURRENT_DATE),
(2, 2, CURRENT_DATE),
(3, 3, CURRENT_DATE),
(4, 1, CURRENT_DATE);

-- LIst of students enrolled in a specified course
SELECT s.student_name
FROM Student s
JOIN Enrollment e ON s.student_id = e.student_id
WHERE e.course_id = (SELECT course_id FROM Course WHERE course_title = 'Introduction to Programming');

-- list of coursers taught by a spescific professor
SELECT c.course_title
FROM Course c
JOIN TeachingAssignment ta ON c.course_id = ta.course_id
WHERE ta.professor_id = (SELECT professor_id FROM Professor WHERE professor_name = 'Prof. Budi');

-- List of prefessors teaching a spescific courses
SELECT p.professor_name
FROM Professor p
JOIN TeachingAssignment ta ON p.professor_id = ta.professor_id
WHERE ta.course_id = (SELECT course_id FROM Course WHERE course_title = 'Golang');

