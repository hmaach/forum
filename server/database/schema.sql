CREATE TABLE IF NOT EXISTS sessions (
    id BIGINT PRIMARY KEY AUTOINCREMENT,
    user_id BIGINT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    email TEXT,
    password TEXT
);
CREATE TABLE IF NOT EXISTS post_category (
    id BIGINT PRIMARY KEY AUTOINCREMENT,
    post_id BIGINT,
    category_id BIGINT,
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
CREATE TABLE IF NOT EXISTS categories (
    id BIGINT PRIMARY KEY AUTOINCREMENT,
    label TEXT
);
CREATE TABLE IF NOT EXISTS comments_reactions (
    id BIGINT PRIMARY KEY AUTOINCREMENT,
    user_id BIGINT,
    comment_id BIGINT,
    type TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (comment_id) REFERENCES comments(id)
);
CREATE TABLE IF NOT EXISTS posts_reactions (
    id BIGINT PRIMARY KEY AUTOINCREMENT,
    user_id BIGINT,
    post_id BIGINT,
    type TEXT,
    created_at TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (post_id) REFERENCES posts(id)
);
CREATE TABLE IF NOT EXISTS posts (
    id BIGINT PRIMARY KEY AUTOINCREMENT,
    user_id BIGINT,
    content TEXT,
    created_at DATETIME,
    title TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS comments (
    id BIGINT PRIMARY KEY AUTOINCREMENT,
    user_id BIGINT,
    post_id BIGINT,
    content TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (post_id) REFERENCES posts(id)
);