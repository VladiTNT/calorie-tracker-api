CREATE TABLE IF NOT EXISTS meal (
    user_name TEXT NOT NULL,
    meal_name TEXT NOT NULL,
    meal_calories INTEGER,
    added_at TEXT DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_name, added_at),
    FOREIGN KEY(user_name) REFERENCES user(user_name) ON DELETE CASCADE
);