CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id BIGINT,
    created_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    email TEXT,
    password TEXT,
    created_at DATETIME
);

CREATE TABLE IF NOT EXISTS post_category (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id BIGINT,
    category_id BIGINT,
    created_at DATETIME,
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    label TEXT,
    created_at DATETIME
);

CREATE TABLE IF NOT EXISTS comments_reactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id BIGINT,
    comment_id BIGINT,
    type TEXT,
    created_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (comment_id) REFERENCES comments(id)
);

CREATE TABLE IF NOT EXISTS posts_reactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id BIGINT,
    post_id BIGINT,
    type TEXT,
    created_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (post_id) REFERENCES posts(id)
);

CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id BIGINT,
    content TEXT,
    title TEXT,
    created_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id BIGINT,
    post_id BIGINT,
    content TEXT,
    created_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (post_id) REFERENCES posts(id)
);
