-- Insert Users
INSERT INTO users (email, username, password) VALUES
('john.doe@example.com', 'JohnDoe', 'password123'),
('jane.smith@example.com', 'JaneSmith', 'securepass'),
('admin@example.com', 'Admin', 'adminpass');

-- Insert Categories
INSERT INTO categories (label) VALUES
('Technology'),
('Health'),
('Travel'),
('Education'),
('Entertainment');

-- Insert Posts
INSERT INTO posts (user_id, title, content) VALUES
(1, 'Understanding SQLite', 'SQLite is a lightweight database that is easy to use and fast.'),
(2, 'Health Tips for 2024', 'Start your day with a glass of water.'),
(3, 'Top 10 Travel Destinations', 'Explore the most beautiful places in the world.'),
(1, 'Education in the Digital Age', 'Technology is transforming education in unprecedented ways.');

-- Link Posts with Categories
INSERT INTO post_category (post_id, category_id) VALUES
(1, 1), -- Post 1 in Technology
(2, 2), -- Post 2 in Health
(3, 3), -- Post 3 in Travel
(4, 4); -- Post 4 in Education

-- Insert Comments
INSERT INTO comments (user_id, post_id, content) VALUES
(2, 1, 'Great article! Very informative.'),
(1, 3, 'I would love to visit these places someday!'),
(3, 2, 'Health is wealth! Thanks for sharing.');

-- Insert Post Reactions
INSERT INTO post_reactions (user_id, post_id, reaction) VALUES
(2, 1, 'like'),
(1, 3, 'like'),
(3, 2, 'dislike');

-- Insert Comment Reactions
INSERT INTO comment_reactions (user_id, comment_id, reaction) VALUES
(1, 1, 'like'),
(3, 2, 'like'),
(2, 3, 'dislike');