CREATE TABLE IF NOT EXISTS users (
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(100)  NOT NULL UNIQUE,
    fullname      VARCHAR(255)  DEFAULT '',
    email         VARCHAR(255)  NOT NULL UNIQUE,
    password_hash VARCHAR(255)  DEFAULT '',
    is_verified   BOOLEAN       DEFAULT FALSE,
    role          VARCHAR(20)   DEFAULT 'user',
    profile_pic   VARCHAR(500)  DEFAULT '',
    created_at    TIMESTAMP     DEFAULT NOW(),
    updated_at    TIMESTAMP     DEFAULT NOW(),
    deleted_at    TIMESTAMP
);