-- +goose Up
-- user Table
CREATE TABLE user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- project Table
CREATE TABLE project (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    /* created_by INTEGER NOT NULL,*/
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
    /* FOREIGN KEY (created_by) REFERENCES user(id) ON DELETE CASCADE */
);

-- Time Entries Table
CREATE TABLE time_entry (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    /*user_id INTEGER NOT NULL,*/
    project_id INTEGER NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    duration INTEGER GENERATED ALWAYS AS (strftime('%s', end_time) - strftime('%s', start_time)) STORED,
    description TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    /*FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE*/
    FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE time_entry;
DROP TABLE project;
DROP TABLE user;