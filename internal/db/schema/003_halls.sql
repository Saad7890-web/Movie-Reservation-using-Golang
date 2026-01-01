CREATE TABLE IF NOT EXISTS halls (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    capacity INT NOT NULL CHECK (capacity > 0)
);
