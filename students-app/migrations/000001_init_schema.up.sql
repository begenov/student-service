CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    gpa FLOAT,
    courses TEXT[]
);

CREATE TABLE admin (
    id SERIAL PRIMARY KEY, 
    email VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    refresh_token TEXT,
    password_hash TEXT
)