-- Fake Data for Users
INSERT INTO users (username, email, password) VALUES
('john_doe', 'john@example.com', 'password123'),
('jane_doe', 'jane@example.com', 'password456'),
('admin', 'admin@example.com', 'adminpass'),
('alice_smith', 'alice@example.com', 'alicepass'),
('bob_jones', 'bob@example.com', 'bobpass');

-- Fake Data for Categories
INSERT INTO categories (label) VALUES
('Technology'),
('Science'),
('Art'),
('Travel'),
('Health'),
('Education');

-- Fake Data for Posts
INSERT INTO posts (user_id, content, created_at, title) VALUES
(1, 'This is the first post about technology.', '2024-11-01 10:00:00', 'Tech Trends 2024'),
(2, 'Art is a reflection of society.', '2024-11-02 12:00:00', 'Understanding Modern Art'),
(1, 'Exploring the wonders of science.', '2024-11-03 14:00:00', 'Science in Everyday Life'),
(3, 'Traveling broadens the mind.', '2024-11-04 09:00:00', 'The Joy of Travel'),
(4, 'Staying healthy is crucial for a happy life.', '2024-11-05 08:00:00', 'Health Tips 101'),
(5, 'Education is the key to success.', '2024-11-06 11:30:00', 'Learning for Life');

-- Fake Data for Post Categories
INSERT INTO post_category (post_id, category_id) VALUES
(1, 1), -- Tech Trends 2024 in Technology
(2, 3), -- Understanding Modern Art in Art
(3, 2), -- Science in Everyday Life in Science
(4, 4), -- The Joy of Travel in Travel
(5, 5), -- Health Tips 101 in Health
(6, 6); -- Learning for Life in Education

-- Fake Data for Comments
INSERT INTO comments (user_id, post_id, content) VALUES
(2, 1, 'Great insights on technology!'),
(3, 2, 'Art truly is inspiring.'),
(1, 3, 'Science is fascinating!'),
(4, 4, 'Traveling is my passion!'),
(5, 5, 'Health tips are always helpful.'),
(2, 6, 'Education should be accessible to everyone.');

-- Fake Data for Comments Reactions
INSERT INTO comments_reactions (user_id, comment_id, type) VALUES
(1, 1, 'like'),
(3, 2, 'dislike'),
(2, 3, 'like'),
(4, 4, 'dislike'),
(5, 5, 'like'),
(1, 6, 'dislike');

-- Fake Data for Posts Reactions
INSERT INTO posts_reactions (user_id, post_id, type, created_at) VALUES
(2, 1, 'like', '2024-11-01 12:00:00'),
(3, 2, 'dislike', '2024-11-02 14:00:00'),
(1, 3, 'like', '2024-11-03 16:00:00'),
(4, 4, 'dislike', '2024-11-04 10:30:00'),
(5, 5, 'like', '2024-11-05 09:15:00'),
(2, 6, 'dislike', '2024-11-06 12:00:00');

-- Fake Data for Sessions
INSERT INTO sessions (user_id) VALUES
(1),
(2),
(3),
(4),
(5);
