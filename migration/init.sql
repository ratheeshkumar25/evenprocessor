CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    raw_event JSONB NOT NULL,
    processed_event JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
