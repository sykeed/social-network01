CREATE TABLE IF NOT EXISTS followers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    follower_id INTEGER NOT NULL,
    following_id INTEGER NOT NULL,
    accepted INTEGER DEFAULT 0 NOT NULL,
    UNIQUE (follower_id, following_id),
    FOREIGN KEY(follower_id) REFERENCES users(id) ,
    FOREIGN KEY(following_id) REFERENCES users(id)
);