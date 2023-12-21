CREATE TABLE IF NOT EXISTS device_events (
    id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL REFERENCES events (id),
    attribute JSONB NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);


--- device_events_created_at_idx ---
CREATE INDEX device_events_created_at_idx ON device_events (created_at);