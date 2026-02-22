CREATE TYPE appointment_status AS ENUM ('pending', 'confirmed', 'completed', 'cancelled');

CREATE TABLE appointments (
    id SERIAL PRIMARY KEY,
    date_time TIMESTAMPTZ,
    car_name varchar(255),
    status appointment_status,
    user_id INTEGER REFERENCES users (id),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
)