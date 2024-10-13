CREATE TABLE IF NOT EXISTS meta.posts (
    id VARCHAR(255) PRIMARY KEY,
    group_id VARCHAR(255),
    FOREIGN KEY (group_id) REFERENCES meta.groups(id),
    message TEXT NOT NULL,
    created_at TIMESTAMP,
    user_id VARCHAR(255),
    origin VARCHAR(50)
);