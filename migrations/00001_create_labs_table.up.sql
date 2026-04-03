CREATE TABLE IF NOT EXISTS labs (
    id               SERIAL PRIMARY KEY,
    labname          VARCHAR(100) UNIQUE NOT NULL,
    title            VARCHAR(255) NOT NULL,
    difficulty       VARCHAR(50),
    category         VARCHAR(100),
    description      TEXT,
    long_description TEXT,
    hints            TEXT[],
    total_solved     INT DEFAULT 0,
    container_id     VARCHAR(100)
);