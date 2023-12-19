CREATE TABLE IF NOT EXISTS applications (
    id SERIAL PRIMARY KEY,
    username_id SERIAL NOT NULL REFERENCES users (id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    api_key TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);
CREATE INDEX idx_api_key ON applications (api_key);