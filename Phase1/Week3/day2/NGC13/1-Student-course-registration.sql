-- Create the Students table
CREATE TABLE students (
    student_id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    major VARCHAR(100),
    year_of_study INTEGER
);

-- Create the Courses table
CREATE TABLE courses (
    course_id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    instructor VARCHAR(100),
    schedule VARCHAR(100),
    credit_hours INTEGER
);

-- Create the Student_Courses table (Associative Entity)
CREATE TABLE student_courses (
    student_id INTEGER REFERENCES students(student_id),
    course_id INTEGER REFERENCES courses(course_id),
    preferred_schedule VARCHAR(100),
    PRIMARY KEY (student_id, course_id)
);

-- Insert sample students
INSERT INTO students (name, email, major, year_of_study) VALUES
('John Doe', 'john.doe@example.com', 'Computer ', 3),
('Jane Smith', 'jane.smith@example.com', 'Seni Budaya', 2),
('Alice Johnson', 'alice.johnson@example.com', 'IPA', 4),
('Bob Brown', 'bob.brown@example.com', 'IPS', 1),
('Eva Wilson', 'eva.wilson@example.com', 'Matematika', 2);

-- Insert sample courses
INSERT INTO courses (title, instructor, schedule, credit_hours) VALUES
('Computer', 'Budi', 'Monday 10:00-11:30', 3),
('Seni Budaya', 'Indah', 'Tuesday 13:00-14:30', 3),
('IPA', 'Tasya', 'Wednesday 9:00-10:30', 4),
('IPS', 'Andreas', 'Thursday 15:00-16:30', 3),
('Matematika', 'John', 'Friday 11:00-12:30', 4);

-- Register students for courses with preferred schedules
INSERT INTO student_courses (student_id, course_id, preferred_schedule) VALUES
(1, 1, 'Monday 10:00-11:30'),
(1, 3, 'Tuesday 9:00-10:30'),
(2, 2, 'Wednesday 13:00-14:30'),
(3, 3, 'Thursday 9:00-10:30'),
(4, 4, 'Friday 15:00-16:30'),
(5, 5, 'Friday 11:00-12:30');

SELECT students.*
FROM students
JOIN student_courses ON students.student_id = student_courses.student_id
WHERE student_courses.course_id = (SELECT course_id FROM courses WHERE title = 'Computer');

SELECT courses.*
FROM courses
JOIN student_courses ON courses.course_id = student_courses.course_id
WHERE student_courses.student_id = (SELECT student_id FROM students WHERE name = 'John Doe');

SELECT preferred_schedule
FROM student_courses
WHERE student_id = (SELECT student_id FROM students WHERE name = 'John Doe')
AND course_id = (SELECT course_id FROM courses WHERE title = 'Introduction to Programming');



