CREATE TABLE IF NOT EXISTS seats (
    id BIGSERIAL PRIMARY KEY,
    hall_id BIGINT NOT NULL REFERENCES halls(id) ON DELETE CASCADE,
    seat_row TEXT NOT NULL,
    seat_number INT NOT NULL,
    UNIQUE (hall_id, seat_row, seat_number)
);
