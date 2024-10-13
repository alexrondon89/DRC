CREATE TABLE IF NOT EXISTS meta.comments (
    id VARCHAR(255) PRIMARY KEY,
    post_id VARCHAR(255),
    FOREIGN KEY (post_id) REFERENCES meta.posts(id),
    message TEXT NOT NULL,
    created_at TIMESTAMP,
    user_id VARCHAR(255),
    origin VARCHAR(50)
);