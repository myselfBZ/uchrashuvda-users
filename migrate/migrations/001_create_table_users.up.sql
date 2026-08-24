CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(255) NOT NULL,
    loc         geography(Point, 4326),
    last_name VARCHAR(255),
    birth_date DATE NOT NULL,
    picture_path TEXT NOT NULL,
    email        TEXT NOT NULL UNIQUE,
    username     VARCHAR(45) NOT NULL UNIQUE,
    password    VARCHAR(255)
);
