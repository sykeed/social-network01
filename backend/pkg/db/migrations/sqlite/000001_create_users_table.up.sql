CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,          
    password TEXT NOT NULL,
    date_of_birth TEXT NOT NULL,        
    gender TEXT NOT NULL,               
    bio TEXT,
    avatar_path TEXT,                   
    is_public INTEGER DEFAULT 1 NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    last_active TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL
);