CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    gender VARCHAR(10),
    birthday DATE,
    launchpad_id VARCHAR(50),
    destination VARCHAR(50),
    launch_date DATE
);
