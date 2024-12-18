DROP TABLE IF EXISTS Loan CASCADE;
DROP TABLE IF EXISTS BookAuthor CASCADE;
DROP TABLE IF EXISTS Book CASCADE;
DROP TABLE IF EXISTS Author CASCADE;
DROP TABLE IF EXISTS Reader CASCADE;

CREATE TABLE Reader (
    reader_id SERIAL PRIMARY KEY,
    first_name VARCHAR(31) NOT NULL,
    last_name VARCHAR(31) NOT NULL,
    phone_number VARCHAR(20) CHECK (phone_number ~ '^\+375[0-9]{9}$'),
    email VARCHAR(255) CHECK (email ~ '^[^@]+@[^@]+\.[^@]+$')
);

CREATE TABLE Author (
    author_id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL
);

CREATE TABLE Book (
    book_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    publication_year INT CHECK (publication_year > 0),
    genre VARCHAR(100)
);

CREATE TABLE BookAuthor (
    book_id INT,
    author_id INT,
    PRIMARY KEY (book_id, author_id),
    FOREIGN KEY (book_id) REFERENCES Book(book_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (author_id) REFERENCES Author(author_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE Loan (
    loan_id SERIAL PRIMARY KEY,
    book_id INT,
    reader_id INT,
    issue_date DATE NOT NULL DEFAULT CURRENT_DATE,
    return_date DATE,
    FOREIGN KEY (book_id) REFERENCES Book(book_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (reader_id) REFERENCES Reader(reader_id) ON DELETE CASCADE ON UPDATE CASCADE,
    CHECK (return_date > issue_date)
);
