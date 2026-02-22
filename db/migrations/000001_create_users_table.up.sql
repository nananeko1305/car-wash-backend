CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar(255),
    last_name varchar(255),
    email varchar(255),
    password varchar(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT users_email_unique UNIQUE (email)
);