CREATE TYPE event AS ENUM ('user_actions', 'ad_actions', 'social_actions');
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    application_id SERIAL NOT NULL REFERENCES applications (id),
    type_event event NOT NULL,
    attribute JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);
CREATE INDEX idx_created_at ON events (created_at);