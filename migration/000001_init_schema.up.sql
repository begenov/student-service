CREATE TABLE student (
    id SERIAL PRIMARY KEY,
    email VARCHAR(50) NOT NULL,
    password_hash VARCHAR(250) NOT NULL,
    name VARCHAR(50) NOT NULL,
    gpa FLOAT,
    refresh_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    courses TEXT[]
);
