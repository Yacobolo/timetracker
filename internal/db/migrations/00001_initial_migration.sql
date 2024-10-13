-- +goose Up
-- user Table
CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    provider TEXT NOT NULL,  -- e.g., 'google', 'github'
    provider_user_id TEXT NOT NULL, -- unique user ID from the external provider
    email TEXT NOT NULL UNIQUE,
    profile_picture TEXT,  -- optional, store user's profile picture URL
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- project Table
CREATE TABLE project (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    -- created_by INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
    -- FOREIGN KEY (created_by) REFERENCES "user"(id) ON DELETE CASCADE
);

-- Time Entries Table
CREATE TABLE time_entry (
    id SERIAL PRIMARY KEY,
    -- user_id INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    duration INTEGER GENERATED ALWAYS AS (EXTRACT(EPOCH FROM end_time) - EXTRACT(EPOCH FROM start_time)) STORED,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    -- FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
    FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE time_entry;
DROP TABLE project;
DROP TABLE "user";
