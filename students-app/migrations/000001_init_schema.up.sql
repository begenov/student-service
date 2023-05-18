CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    gpa FLOAT,
    courses TEXT[]
);
