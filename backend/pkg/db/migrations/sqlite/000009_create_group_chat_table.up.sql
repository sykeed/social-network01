CREATE TABLE IF NOT EXISTS group_chat (
    id INTEGER PRIMARY KEY AUTOINCREMENT,               
    group_id INTEGER NOT NULL,                     
    sender_id INTEGER NOT NULL,                     
    message TEXT NOT NULL,                           
    created_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(id),
    FOREIGN KEY (sender_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS message_seen (
    id INTEGER PRIMARY KEY AUTOINCREMENT,         
    message_id INTEGER NOT NULL,                   
    user_id INTEGER NOT NULL,                    
    UNIQUE (message_id, user_id),                  
    FOREIGN KEY (message_id) REFERENCES group_chat(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);