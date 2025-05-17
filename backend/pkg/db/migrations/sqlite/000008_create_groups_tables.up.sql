CREATE TABLE IF NOT EXISTS groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    admin_id INTEGER NOT NULL, 
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY(admin_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS group_members (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,           
    group_id INTEGER NOT NULL,           
    status TEXT DEFAULT 'pending' NOT NULL,                                    
    role TEXT DEFAULT 'member' NOT NULL,                                              
    joined_at TEXT,                            
    UNIQUE (user_id, group_id),                
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(group_id) REFERENCES groups(id)
);