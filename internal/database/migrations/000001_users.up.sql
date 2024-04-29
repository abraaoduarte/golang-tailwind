CREATE TABLE IF NOT EXISTS users (
    id SERIAL primary key,
    name VARCHAR(255)  not null,
    email VARCHAR(255) not null UNIQUE,
    birthdate DATE not null,
    is_admin boolean DEFAULT false,
    created_at TIMESTAMP default now()
);
