CREATE TABLE student (
    id SERIAL PRIMARY KEY,
    email VARCHAR(50) NOT NULL,
    password_hash VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    gpa FLOAT,
    refresh_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    courses TEXT[]
);

CREATE TABLE admin (
    id SERIAL PRIMARY KEY, 
    email VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    refresh_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    password_hash TEXT
)