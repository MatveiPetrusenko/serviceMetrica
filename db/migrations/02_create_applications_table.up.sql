CREATE TABLE IF NOT EXISTS applications (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users (id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    api_key TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

--- applications_api_key_idx ---
CREATE UNIQUE INDEX applications_api_key_idx ON applications (api_key);