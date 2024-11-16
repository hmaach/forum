-- Fake Data for Users (25 users)
INSERT INTO users (username, email, password) VALUES
('john_doe', 'john@example.com', 'password123'),
('jane_smith', 'jane@example.com', 'password456'),
('admin_user', 'admin@example.com', 'adminpass'),
('alice_wonder', 'alice@example.com', 'alicepass'),
('bob_builder', 'bob@example.com', 'bobpass'),
('sarah_connor', 'sarah@example.com', 'sarahpass'),
('mike_jones', 'mike@example.com', 'mikepass'),
('emma_watson', 'emma@example.com', 'emmapass'),
('david_miller', 'david@example.com', 'davidpass'),
('lisa_brown', 'lisa@example.com', 'lisapass'),
('tech_guru', 'guru@example.com', 'gurupass'),
('travel_lover', 'travel@example.com', 'travelpass'),
('health_expert', 'health@example.com', 'healthpass'),
('art_critic', 'art@example.com', 'artpass'),
('science_nerd', 'science@example.com', 'sciencepass'),
('fitness_freak', 'fitness@example.com', 'fitnesspass'),
('food_lover', 'food@example.com', 'foodpass'),
('book_worm', 'books@example.com', 'bookpass'),
('music_fan', 'music@example.com', 'musicpass'),
('movie_buff', 'movies@example.com', 'moviepass'),
('nature_explorer', 'nature@example.com', 'naturepass'),
('tech_reviewer', 'reviewer@example.com', 'reviewpass'),
('daily_blogger', 'blogger@example.com', 'blogpass'),
('photo_pro', 'photo@example.com', 'photopass'),
('gaming_master', 'gaming@example.com', 'gamepass');

-- Expanded Categories (12 categories)
INSERT INTO categories (label) VALUES
('Technology'),
('Science'),
('Art'),
('Travel'),
('Health'),
('Education'),
('Food'),
('Sports'),
('Entertainment'),
('Business'),
('Lifestyle'),
('Gaming');

-- Fake Data for Posts (20 posts)
INSERT INTO posts (user_id, content, created_at, title) VALUES
(1, 'Exploring the latest trends in AI and machine learning. The rapid advancement of artificial intelligence is reshaping various industries, from healthcare to finance. Here''s an in-depth look at how AI is transforming our daily lives.', '2024-01-15 10:00:00', 'AI Revolution 2024'),
(11, 'Review of the newest smartphone releases and their groundbreaking features. From improved cameras to revolutionary processors, these devices are pushing the boundaries of mobile technology.', '2024-02-01 14:30:00', 'Smartphone Innovation'),
(22, 'The future of cloud computing and its impact on business operations. Cloud technology continues to evolve, offering more scalable and efficient solutions for enterprises of all sizes.', '2024-02-15 09:15:00', 'Cloud Tech Trends'),
(15, 'Recent breakthroughs in quantum computing and their potential applications. Scientists have achieved new milestones in quantum supremacy, opening doors to unprecedented computational capabilities.', '2024-02-20 11:00:00', 'Quantum Computing Advances'),
(3, 'Understanding climate change: Latest research and global impact. New studies reveal concerning trends in global temperature rises and their effects on ecosystems worldwide.', '2024-03-01 13:45:00', 'Climate Science Update'),
(8, 'Discoveries in space exploration: What we learned from recent missions. The latest Mars rover findings and observations from deep space telescopes are reshaping our understanding of the cosmos.', '2024-03-10 16:20:00', 'Space Exploration 2024'),
(14, 'Contemporary art movements shaping modern culture. How digital art and NFTs are revolutionizing the art world and creating new opportunities for artists.', '2024-03-15 10:30:00', 'Modern Art Trends'),
(2, 'Street art: A powerful medium for social commentary. Urban artists are using public spaces to address pressing social issues and spark important conversations.', '2024-03-20 14:15:00', 'Street Art Revolution'),
(12, 'Hidden gems: Exploring off-the-beaten-path destinations. Discover these lesser-known locations that offer unique experiences for adventurous travelers.', '2024-03-25 09:45:00', 'Secret Travel Spots'),
(4, 'Sustainable tourism: How to travel responsibly in 2024. Tips and guidelines for minimizing your environmental impact while exploring the world.', '2024-04-01 11:30:00', 'Eco-Friendly Travel'),
(5, 'The importance of mental health awareness in modern society. Understanding and addressing mental health issues has become increasingly crucial in our fast-paced world.', '2024-04-05 13:20:00', 'Mental Health Matters'),
(6, 'Latest developments in renewable energy technology. Solar and wind power innovations are making sustainable energy more accessible than ever.', '2024-04-10 15:45:00', 'Green Energy Future'),
(7, 'The evolution of remote work culture. How companies and employees are adapting to the new normal of hybrid work environments.', '2024-04-15 09:30:00', 'Remote Work Revolution'),
(8, 'Emerging trends in cybersecurity. Protecting digital assets has become more challenging as cyber threats continue to evolve.', '2024-04-20 14:20:00', 'Cybersecurity Trends'),
(9, 'The impact of social media on modern society. Examining both the benefits and drawbacks of our increasingly connected world.', '2024-04-25 11:15:00', 'Social Media Impact'),
(10, 'Sustainable fashion: The future of the clothing industry. How eco-friendly practices are reshaping fashion consumption.', '2024-04-30 16:40:00', 'Sustainable Fashion'),
(11, 'The rise of plant-based diets and their environmental impact. More people are choosing plant-based options for health and environmental reasons.', '2024-05-05 10:25:00', 'Plant-Based Revolution'),
(12, 'Virtual reality in education: New ways of learning. How VR technology is transforming the educational experience.', '2024-05-10 13:50:00', 'VR in Education'),
(13, 'The future of autonomous vehicles. Self-driving technology is advancing rapidly, but challenges remain.', '2024-05-15 15:30:00', 'Autonomous Driving'),
(14, 'Digital privacy in the age of big data. Protecting personal information has become increasingly important in our connected world.', '2024-05-20 12:10:00', 'Digital Privacy');

-- Post Categories
INSERT INTO post_category (post_id, category_id) VALUES
(1, 1), (1, 2),
(1, 6),(2, 1),
(2, 10),(3, 1),
(3, 10), (3, 6),
(4, 1), (4, 2),
(5, 2), (5, 5),
(6, 2), (6, 1),
(7, 3), (7, 9),
(8, 3), (8, 11),
(9, 4), (9, 11),
(10, 4), (10, 5),
(11, 5), (11, 11),
(12, 1), (12, 2),
(13, 10), (13, 11),
(14, 1), (14, 6),
(15, 9), (15, 11),
(16, 11), (16, 10),
(17, 5), (17, 7),
(18, 6), (18, 1),
(19, 1), (19, 10),
(20, 1), (20, 11);

-- Comments
INSERT INTO comments (user_id, post_id, content) VALUES
(2, 1, 'Fascinating insights into AI development. The potential applications in healthcare are particularly promising.'),
(5, 1, 'Great article! Would love to see more specific examples of AI implementation in different industries.'),
(8, 1, 'The ethical implications of AI advancement need more discussion. How do we ensure responsible development?'),
(3, 2, 'These new smartphones are incredible. The camera capabilities are mind-blowing.'),
(6, 2, 'Interesting comparison of different models. Very helpful for making a purchase decision.'),
(9, 3, 'Cloud computing has revolutionized how we do business. Great analysis!'),
(4, 4, 'Quantum computing is the future. Can''t wait to see what developments come next.'),
(7, 5, 'Climate change is such a critical issue. We need more awareness like this.'),
(10, 6, 'Space exploration continues to amaze me. Excellent coverage of recent missions.');

-- Posts Reactions
INSERT INTO posts_reactions (user_id, post_id, type, created_at) VALUES
(2, 1, 'like', '2024-01-15 10:30:00'),
(3, 1, 'like', '2024-01-15 11:15:00'),
(4, 1, 'like', '2024-01-15 12:00:00'),
(5, 1, 'dislike', '2024-01-15 13:45:00'),
(6, 2, 'like', '2024-02-01 15:00:00'),
(7, 2, 'like', '2024-02-01 16:30:00'),
(8, 3, 'like', '2024-02-15 10:00:00'),
(9, 3, 'dislike', '2024-02-15 11:30:00'),
(10, 4, 'like', '2024-02-20 12:00:00');

-- Comments Reactions
INSERT INTO comments_reactions (user_id, comment_id, type) VALUES
(1, 1, 'like'),
(3, 1, 'like'),
(4, 1, 'like'),
(5, 1, 'dislike'),
(2, 2, 'like'),
(4, 2, 'like'),
(6, 3, 'like'),
(7, 3, 'dislike'),
(8, 4, 'like');

-- Active Sessions
INSERT INTO sessions (user_id)
SELECT id FROM users WHERE id <= 15;