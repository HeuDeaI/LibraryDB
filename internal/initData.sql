INSERT INTO Reader (first_name, last_name, phone_number, email) VALUES
('John', 'Doe', '+375291234567', 'john.doe@example.com'),
('Jane', 'Smith', '+375291234568', 'jane.smith@example.com'),
('Alice', 'Johnson', '+375291234569', 'alice.johnson@example.com'),
('Bob', 'Brown', '+375291234570', 'bob.brown@example.com'),
('Charlie', 'Davis', '+375291234571', 'charlie.davis@example.com'),
('Diana', 'Miller', '+375291234572', 'diana.miller@example.com'),
('Eva', 'Wilson', '+375291234573', 'eva.wilson@example.com'),
('Frank', 'Taylor', '+375291234574', 'frank.taylor@example.com'),
('Grace', 'Anderson', '+375291234575', 'grace.anderson@example.com'),
('Hank', 'Thomas', '+375291234576', 'hank.thomas@example.com'),
('Ivy', 'Jackson', '+375291234577', 'ivy.jackson@example.com'),
('Jack', 'White', '+375291234578', 'jack.white@example.com'),
('Karen', 'Harris', '+375291234579', 'karen.harris@example.com'),
('Larry', 'Martin', '+375291234580', 'larry.martin@example.com'),
('Megan', 'Thompson', '+375291234581', 'megan.thompson@example.com'),
('Nathan', 'Garcia', '+375291234582', 'nathan.garcia@example.com'),
('Olivia', 'Martinez', '+375291234583', 'olivia.martinez@example.com'),
('Paul', 'Robinson', '+375291234584', 'paul.robinson@example.com'),
('Quinn', 'Clark', '+375291234585', 'quinn.clark@example.com'),
('Rita', 'Rodriguez', '+375291234586', 'rita.rodriguez@example.com');

INSERT INTO Author (first_name, last_name) VALUES
('George', 'Orwell'),
('J.K.', 'Rowling'),
('J.R.R.', 'Tolkien'),
('Agatha', 'Christie'),
('Mark', 'Twain'),
('Ernest', 'Hemingway'),
('F. Scott', 'Fitzgerald'),
('Leo', 'Tolstoy'),
('Virginia', 'Woolf'),
('Herman', 'Melville'),
('Jane', 'Austen'),
('Charles', 'Dickens'),
('Gabriel', 'Garcia Marquez'),
('Harper', 'Lee'),
('Stephen', 'King'),
('Oscar', 'Wilde'),
('Ray', 'Bradbury'),
('Toni', 'Morrison'),
('John', 'Steinbeck'),
('Emily', 'Dickinson');

INSERT INTO Book (title, publication_year, genre) VALUES
('1984', 1949, 'Dystopian'),
('Harry Potter and the Philosophers Stone', 1997, 'Fantasy'),
('The Hobbit', 1937, 'Fantasy'),
('Murder on the Orient Express', 1934, 'Mystery'),
('Adventures of Huckleberry Finn', 1884, 'Adventure'),
('The Old Man and the Sea', 1952, 'Fiction'),
('The Great Gatsby', 1925, 'Fiction'),
('War and Peace', 1869, 'Historical Fiction'),
('To the Lighthouse', 1927, 'Modernist Fiction'),
('Moby Dick', 1851, 'Adventure'),
('Pride and Prejudice', 1813, 'Romance'),
('A Tale of Two Cities', 1859, 'Historical Fiction'),
('One Hundred Years of Solitude', 1967, 'Magic Realism'),
('To Kill a Mockingbird', 1960, 'Fiction'),
('The Shining', 1977, 'Horror'),
('The Picture of Dorian Gray', 1890, 'Philosophical Fiction'),
('Fahrenheit 451', 1953, 'Dystopian'),
('Beloved', 1987, 'Historical Fiction'),
('The Grapes of Wrath', 1939, 'Historical Fiction'),
('The Color Purple', 1982, 'Fiction'),
('The Martian', 2011, 'Science Fiction');


INSERT INTO BookAuthor (book_id, author_id) VALUES
(1, 1),  -- 1984 by George Orwell
(2, 2),  -- Harry Potter and the Philosopher's Stone by J.K. Rowling
(3, 3),  -- The Hobbit by J.R.R. Tolkien
(4, 4),  -- Murder on the Orient Express by Agatha Christie
(5, 5),  -- Adventures of Huckleberry Finn by Mark Twain
(6, 6),  -- The Old Man and the Sea by Ernest Hemingway
(7, 7),  -- The Great Gatsby by F. Scott Fitzgerald
(8, 8),  -- War and Peace by Leo Tolstoy
(9, 9),  -- To the Lighthouse by Virginia Woolf
(10, 10), -- Moby Dick by Herman Melville
(11, 11), -- Pride and Prejudice by Jane Austen
(12, 12), -- A Tale of Two Cities by Charles Dickens
(13, 13), -- One Hundred Years of Solitude by Gabriel Garcia Marquez
(14, 14), -- To Kill a Mockingbird by Harper Lee
(15, 15), -- The Shining by Stephen King
(16, 16), -- The Picture of Dorian Gray by Oscar Wilde
(17, 17), -- Fahrenheit 451 by Ray Bradbury
(18, 18), -- Beloved by Toni Morrison
(19, 19), -- The Grapes of Wrath by John Steinbeck
(20, 20); -- The Martian by Andy Weir


INSERT INTO Loan (book_id, reader_id, issue_date, return_date) VALUES
(1, 1, '2023-02-10', '2023-02-17'),
(2, 2, '2023-01-12', '2023-01-19'),
(3, 3, '2023-01-14', '2023-01-21'),
(4, 4, '2023-01-15', '2023-01-22'),
(5, 5, '2023-01-16', '2023-01-23'),
(6, 6, '2023-01-17', '2023-01-24'),
(7, 7, '2023-01-18', '2023-01-25'),
(8, 8, '2023-01-19', '2023-01-26'),
(9, 9, '2023-01-20', '2023-01-29'), 
(10, 10, '2023-01-21', '2023-01-27'),
(11, 11, '2023-01-22', '2023-02-01'), 
(12, 12, '2023-01-23', '2023-01-31'),
(13, 13, '2023-01-24', '2023-02-02'),
(14, 14, '2023-01-25', '2023-02-03'),
(15, 15, '2023-01-26', '2023-02-04'),
(16, 16, '2023-01-27', '2023-02-02'), 
(17, 17, '2023-01-28', '2023-02-05'), 
(18, 18, '2023-01-29', '2023-02-06'), 
(19, 19, '2023-01-30', '2023-02-07'), 
(20, 20, '2023-01-31', '2023-02-08'); 
