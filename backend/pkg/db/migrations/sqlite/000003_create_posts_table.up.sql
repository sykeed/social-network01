CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    group_id INTEGER,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    image TEXT,
    created_at TEXT NOT NULL,
    privacy TEXT,
    can_see TEXT,
        CHECK (
        (group_id IS NULL AND privacy IS NOT NULL AND privacy IN ('public', 'private', 'semi-private')) OR
        (group_id IS NOT NULL AND privacy IS NULL)
    ),
    foreign key(group_id) references groups(id),
    foreign key(user_id) references users(id)
);

CREATE TABLE IF NOT EXISTS private_posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER NOT NULL,
    can_see INTEGER NOT NULL,
    foreign key(post_id) references posts(id),
    foreign key(can_see) references users(id)
);