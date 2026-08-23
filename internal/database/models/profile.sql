CREATE TABLE IF NOT EXISTS user_profile (
    user_name TEXT NOT NULL PRIMARY KEY UNIQUE,
    target_calories INTEGER,
    FOREIGN KEY(user_name) REFERENCES user(user_name) ON DELETE CASCADE
);