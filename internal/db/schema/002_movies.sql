CREATE TABLE IF NOT EXISTS movies (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    poster_url TEXT,
    genre TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
