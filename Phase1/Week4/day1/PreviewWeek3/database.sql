CREATE TABLE bookstore (
    order_id INT,
    customer_name VARCHAR(255),
    customer_email VARCHAR(255),
    book_title VARCHAR(255),
    book_type VARCHAR(255), -- E-Book or Physical
    author_name VARCHAR(255),
    author_email VARCHAR(255),
    order_date DATE,
    price DECIMAL(10, 2)
);

INSERT INTO bookstore VALUES
(1, 'John Doe', 'john.doe@example.com', 'Database Design', 'Physical', 'Jane Smith', 'jane.smith@example.com', '2023-08-10', 25.99),
(2, 'John Doe', 'john.doe@example.com', 'Web Development', 'E-Book', 'Tom Brown', 'tom.brown@example.com', '2023-08-11', 19.99),
(3, 'Alice Bob', 'alice.bob@example.com', 'Database Design', 'E-Book', 'Jane Smith', 'jane.smith@example.com', '2023-08-12', 20.99);

-- Normalisasi
CREATE TABLE authors (
    author_id INT PRIMARY KEY AUTO_INCREMENT,
    author_name VARCHAR(255),
    author_email VARCHAR(255)
);

CREATE TABLE customers (
    customer_id INT PRIMARY KEY AUTO_INCREMENT,
    customer_name VARCHAR(255),
    customer_email VARCHAR(255)
);

CREATE TABLE books (
    book_id INT PRIMARY KEY AUTO_INCREMENT,
    book_title VARCHAR(255),
    book_type VARCHAR(255) -- E-Book or Physical
);

CREATE TABLE orders (
    order_id INT PRIMARY KEY AUTO_INCREMENT,
    customer_id INT,
    book_id INT,
    author_id INT,
    order_date DATE,
    price DECIMAL(10, 2),
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id),
    FOREIGN KEY (book_id) REFERENCES books(book_id),
    FOREIGN KEY (author_id) REFERENCES authors(author_id)
);
