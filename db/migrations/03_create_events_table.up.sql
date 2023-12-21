CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    application_id INTEGER NOT NULL REFERENCES applications (id),
    type_event TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

--- events_type_event_idx ---
CREATE INDEX events_type_event_idx ON events (type_event);