-- Insert categories (if not already done)
INSERT OR IGNORE INTO categories (label) VALUES
('Technology'),
('Health'),
('Travel'),
('Education'),
('Entertainment');

-- Insert 50 posts with associated users (1 to 5, as example users)
INSERT INTO posts (user_id, title, content) VALUES
(1, 'The Future of AI Technology', 'AI is shaping the future in many exciting ways. In this article, we explore the potential advancements in machine learning and robotics.'),
(2, 'Understanding Blockchain and Cryptocurrencies', 'Blockchain technology is revolutionizing industries. This post discusses how blockchain works and its applications beyond cryptocurrencies.'),
(3, 'Top 10 Travel Destinations for 2025', 'Looking for travel ideas? We’ve compiled the top 10 travel destinations for 2025, with options for every type of traveler.'),
(4, 'Mental Health in the Digital Age', 'The rise of social media has impacted mental health in profound ways. This article explores the challenges and solutions for mental well-being online.'),
(5, 'The Importance of Education in the 21st Century', 'Education is the key to unlocking potential. We look at how modern education systems are evolving to meet the needs of today’s world.'),
(1, 'Exploring the World of Renewable Energy', 'Renewable energy is transforming our global energy landscape. Learn about the latest advancements in solar, wind, and hydropower.'),
(2, 'The Impact of 5G on Connectivity', '5G promises faster speeds and better connectivity. Here’s what you need to know about the upcoming 5G revolution and its global impact.'),
(3, 'How to Start Your Own Business', 'Starting a business requires more than just an idea. We discuss the essential steps for turning your entrepreneurial dream into reality.'),
(4, 'Top 5 Health Trends in 2025', 'Health trends are constantly changing. This post highlights the top health trends for 2025, from fitness routines to diet innovations.'),
(5, 'The Rise of Virtual Reality in Education', 'Virtual reality is transforming how we learn. This article explores how VR is being used in education to enhance the learning experience.'),
(1, 'The Ethics of Artificial Intelligence', 'AI ethics is a rapidly growing field. This post delves into the ethical implications of AI, from bias to privacy concerns.'),
(2, 'How to Travel on a Budget', 'Traveling doesn’t have to break the bank. In this guide, we share tips for traveling affordably without compromising on experience.'),
(3, 'Why Sleep is Essential for Mental Health', 'A good night’s sleep is vital for mental health. Learn about the science behind sleep and how it impacts our emotional and physical well-being.'),
(4, 'The Evolution of Mobile Phones', 'Mobile phones have come a long way since their invention. This article traces the evolution of mobile phones and their impact on society.'),
(5, 'Technology’s Role in Healthcare Innovation', 'Technology is changing healthcare for the better. This post examines the innovative technologies improving patient care and treatment.'),
(1, 'How Social Media Influences Our Lives', 'Social media plays a significant role in shaping our lives. We explore the positive and negative effects of social media on individuals and society.'),
(2, 'Exploring the Internet of Things (IoT)', 'The Internet of Things (IoT) is connecting everything. This post explains what IoT is and how it’s transforming industries like healthcare, transportation, and home automation.'),
(3, 'Sustainable Travel: A Guide to Eco-Friendly Adventures', 'Sustainable travel is becoming more popular. Learn how to reduce your carbon footprint while exploring the world.'),
(4, 'How Technology is Revolutionizing the Healthcare Industry', 'From telemedicine to AI diagnostics, technology is making healthcare more accessible and efficient. This article highlights key innovations.'),
(5, 'The Best Online Learning Platforms of 2025', 'Online learning is more accessible than ever. Discover the top platforms for enhancing your skills and gaining new knowledge in 2025.'),
(1, 'AI and Ethics: Striking a Balance', 'Artificial intelligence presents many ethical dilemmas. This post examines how we can balance innovation with moral responsibility in AI development.'),
(2, 'The Future of Autonomous Vehicles', 'Autonomous vehicles are on the horizon. In this article, we explore how self-driving cars could change the future of transportation.'),
(3, 'Tips for Effective Time Management', 'Managing time effectively is crucial for success. This post shares actionable strategies to help you maximize your productivity.'),
(4, 'The Benefits of a Plant-Based Diet', 'A plant-based diet has numerous health benefits. Learn how switching to plant-based foods can improve your overall well-being.'),
(5, 'Advancements in Space Exploration', 'Space exploration has made tremendous strides. This post covers the latest achievements in space travel and the search for life beyond Earth.'),
(1, 'The Impact of Artificial Intelligence on Job Markets', 'AI is changing industries and jobs. This post discusses the impact of AI on employment and the skills needed for the future workforce.'),
(2, 'How to Stay Fit While Traveling', 'Staying fit while traveling can be challenging. Here are some simple tips for keeping active while on the go.'),
(3, 'How To Build A Personal Brand Online', 'In the digital age, personal branding is essential. This post gives you a roadmap for creating a strong personal brand that stands out online.'),
(4, 'The Role of Technology in Combating Climate Change', 'Technology is a powerful tool in the fight against climate change. This article examines the innovative ways technology is being used to reduce emissions and protect the planet.'),
(5, 'What Is Deep Learning and How Does It Work?', 'Deep learning is an exciting field in AI. This post explains what deep learning is and how it’s used in areas like natural language processing and computer vision.'),
(1, 'Artificial Intelligence in Healthcare', 'AI is revolutionizing healthcare in many ways, from diagnostics to treatment. This post explores the future of AI in the medical field.'),
(2, 'The Benefits of Traveling Solo', 'Solo travel offers a unique experience. In this post, we discuss the benefits of traveling alone and how it can be a life-changing experience.'),
(3, 'Healthy Habits for a Productive Workday', 'Building healthy habits is key to staying productive at work. We share tips on maintaining a balanced lifestyle for professional success.'),
(4, 'Blockchain Beyond Bitcoin', 'Blockchain is not just for cryptocurrencies. This post discusses how blockchain technology is being used in supply chains, healthcare, and beyond.'),
(5, 'Breaking Down the Latest Trends in Fitness', 'Fitness is always evolving. This post covers the latest trends in fitness, from virtual workouts to biohacking for peak performance.'),
(1, 'The Importance of Data Privacy', 'Data privacy is becoming more important than ever. In this article, we discuss why data privacy matters and how to protect your personal information online.'),
(2, 'How to Maintain a Healthy Work-Life Balance', 'Maintaining work-life balance is essential for long-term well-being. This post shares practical tips for finding harmony between your career and personal life.'),
(3, 'The Impact of Technology on Education', 'Technology is reshaping education. This post highlights the role of technology in enhancing learning and accessibility in schools and universities.'),
(4, 'How to Manage Stress Effectively', 'Stress can affect our health and well-being. This article offers strategies for managing stress in a healthy and productive way.'),
(5, 'The Rise of E-Sports: A New Era of Gaming', 'E-sports has exploded in popularity. Learn about the rise of competitive gaming and its global impact on the entertainment industry.');

-- Insert categories for each post (post_id should be consistent with the posts table)
INSERT INTO post_category (post_id, category_id) VALUES
(1, 1), (2, 2), (3, 3), (4, 4), (5, 5),
(6, 1), (7, 2), (8, 3), (9, 4), (10, 5),
(11, 1), (12, 2), (13, 3), (14, 4), (15, 5),
(16, 1), (17, 2), (18, 3), (19, 4), (20, 5),
(21, 1), (22, 2), (23, 3), (24, 4), (25, 5),
(26, 1), (27, 2), (28, 3), (29, 4), (30, 5),
(31, 1), (32, 2), (33, 3), (34, 4), (35, 5),
(36, 1), (37, 2), (38, 3), (39, 4), (40, 5),
(41, 1), (42, 2), (43, 3), (44, 4), (45, 5),
(46, 1), (47, 2), (48, 3), (49, 4), (50, 5);

-- Insert reactions for posts (each user can react to a post)
INSERT INTO post_reactions (user_id, post_id, reaction) VALUES
(1, 1, 'like'), (2, 2, 'like'), (3, 3, 'dislike'), (4, 4, 'like'), (5, 5, 'dislike'),
(1, 6, 'like'), (2, 7, 'dislike'), (3, 8, 'like'), (4, 9, 'dislike'), (5, 10, 'like'),
-- Add reactions for posts 11 to 50 (following the same pattern above)

-- Insert comments for random posts
INSERT INTO comments (user_id, post_id, content) VALUES
(1, 1, 'Great article on AI technology!'), (2, 2, 'I love this post on blockchain.'),
(3, 3, 'Can’t wait to visit these travel destinations!'), (4, 4, 'Such a relevant topic, mental health matters!'),
(5, 5, 'Great insights on education in today’s world.'), -- Continue for other posts

-- Insert reactions for comments (user_id, comment_id, reaction)
INSERT INTO comment_reactions (user_id, comment_id, reaction) VALUES
(1, 1, 'like'), (2, 2, 'dislike'), (3, 3, 'like'), (4, 4, 'dislike'), (5, 5, 'like');
-- Continue for other comments
