CREATE TABLE sessions (
  id BIGINT AUTOINCREMENT PRIMARY KEY,
  user_id BIGINT,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE users (
  id BIGINT AUTOINCREMENT PRIMARY KEY,
  username TEXT,
  email TEXT,
  password TEXT
);

CREATE TABLE post_category (
  id BIGINT AUTOINCREMENT PRIMARY KEY,
  post_id BIGINT,
  category_id BIGINT,
  FOREIGN KEY (post_id) REFERENCES posts(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE categories (
  id BIGINT AUTOINCREMENT PRIMARY KEY,
  label TEXT
);

CREATE TABLE comments_reactions (
  id BIGINT AUTOINCREMENT PRIMARY KEY,
  user_id BIGINT,
  comment_id BIGINT,
  type TEXT,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (comment_id) REFERENCES comments(id)
);

CREATE TABLE posts_reactions (
  id BIGINT AUTOINCREMENT PRIMARY KEY,
  user_id BIGINT,
  post_id BIGINT,
  type TEXT,
  created_at TEXT,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (post_id) REFERENCES posts(id)
);

CREATE TABLE posts (
  id BIGINT AUTOINCREMENT PRIMARY KEY,
  user_id BIGINT,
  content TEXT,
  created_at DATETIME,
  title TEXT,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE comments (
  id BIGINT AUTOINCREMENT PRIMARY KEY,
  user_id BIGINT,
  post_id BIGINT,
  content TEXT,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (post_id) REFERENCES posts(id)
);