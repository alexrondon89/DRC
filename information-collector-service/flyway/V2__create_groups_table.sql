CREATE TABLE IF NOT EXISTS meta.groups(
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    privacy VARCHAR(50),
    origin VARCHAR(50),
    updated_time TIMESTAMP
);