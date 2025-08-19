-- =========================
-- Users (5)
-- =========================
INSERT INTO `user` (id, created_at, updated_at, username, post_num) VALUES
                                                                        (1, NOW(), NOW(), 'alice',   7),
                                                                        (2, NOW(), NOW(), 'bob',     6),
                                                                        (3, NOW(), NOW(), 'charlie', 5),
                                                                        (4, NOW(), NOW(), 'david',   4),
                                                                        (5, NOW(), NOW(), 'eva',     3);

-- =========================
-- Posts (25) - uneven per user
-- Alice (7): posts 1..7
-- Bob   (6): posts 8..13
-- Charlie(5): posts 14..18
-- David (4): posts 19..22
-- Eva   (3): posts 23..25
-- =========================

-- Alice (1..7) - 4,6 无评论
INSERT INTO `post` (id, created_at, updated_at, title, body, user_id, comment_state) VALUES
                                                                                         (1, NOW(), NOW(), 'Alice Post 1', 'Content A1', 1, '有评论'),
                                                                                         (2, NOW(), NOW(), 'Alice Post 2', 'Content A2', 1, '有评论'),
                                                                                         (3, NOW(), NOW(), 'Alice Post 3', 'Content A3', 1, '有评论'),
                                                                                         (4, NOW(), NOW(), 'Alice Post 4', 'Content A4', 1, '无评论'),
                                                                                         (5, NOW(), NOW(), 'Alice Post 5', 'Content A5', 1, '有评论'),
                                                                                         (6, NOW(), NOW(), 'Alice Post 6', 'Content A6', 1, '无评论'),
                                                                                         (7, NOW(), NOW(), 'Alice Post 7', 'Content A7', 1, '有评论');

-- Bob (8..13) - 10,13 无评论
INSERT INTO `post` (id, created_at, updated_at, title, body, user_id, comment_state) VALUES
                                                                                         (8,  NOW(), NOW(), 'Bob Post 1', 'Content B1', 2, '有评论'),
                                                                                         (9,  NOW(), NOW(), 'Bob Post 2', 'Content B2', 2, '有评论'),
                                                                                         (10, NOW(), NOW(), 'Bob Post 3', 'Content B3', 2, '无评论'),
                                                                                         (11, NOW(), NOW(), 'Bob Post 4', 'Content B4', 2, '有评论'),
                                                                                         (12, NOW(), NOW(), 'Bob Post 5', 'Content B5', 2, '有评论'),
                                                                                         (13, NOW(), NOW(), 'Bob Post 6', 'Content B6', 2, '无评论');

-- Charlie (14..18) - 17,18 无评论
INSERT INTO `post` (id, created_at, updated_at, title, body, user_id, comment_state) VALUES
                                                                                         (14, NOW(), NOW(), 'Charlie Post 1', 'Content C1', 3, '有评论'),
                                                                                         (15, NOW(), NOW(), 'Charlie Post 2', 'Content C2', 3, '有评论'),
                                                                                         (16, NOW(), NOW(), 'Charlie Post 3', 'Content C3', 3, '有评论'),
                                                                                         (17, NOW(), NOW(), 'Charlie Post 4', 'Content C4', 3, '无评论'),
                                                                                         (18, NOW(), NOW(), 'Charlie Post 5', 'Content C5', 3, '无评论');

-- David (19..22) - 21 无评论
INSERT INTO `post` (id, created_at, updated_at, title, body, user_id, comment_state) VALUES
                                                                                         (19, NOW(), NOW(), 'David Post 1', 'Content D1', 4, '有评论'),
                                                                                         (20, NOW(), NOW(), 'David Post 2', 'Content D2', 4, '有评论'),
                                                                                         (21, NOW(), NOW(), 'David Post 3', 'Content D3', 4, '无评论'),
                                                                                         (22, NOW(), NOW(), 'David Post 4', 'Content D4', 4, '有评论');

-- Eva (23..25) - 24 无评论
INSERT INTO `post` (id, created_at, updated_at, title, body, user_id, comment_state) VALUES
                                                                                         (23, NOW(), NOW(), 'Eva Post 1', 'Content E1', 5, '有评论'),
                                                                                         (24, NOW(), NOW(), 'Eva Post 2', 'Content E2', 5, '无评论'),
                                                                                         (25, NOW(), NOW(), 'Eva Post 3', 'Content E3', 5, '有评论');

-- =========================
-- Comments (80 total, uneven)
-- 分配方案（仅给“有评论”的文章）：
-- Alice: 1:12, 2:6, 3:3, 5:5, 7:2  => 28
-- Bob:   8:10, 9:3, 11:6, 12:4      => 23
-- Charlie: 14:8, 15:4, 16:2         => 14
-- David: 19:5, 20:4, 22:3           => 12
-- Eva:   23:2, 25:1                 => 3
-- 合计：80
-- =========================

-- Helper:
-- publish_user_id 从 1..5 里挑，尽量避免和作者相同（不是硬性要求）

-- Post 1 (Alice) - 12
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Great post!', 1, 2),
                                                                                          (NOW(), NOW(), 'I like it!', 1, 3),
                                                                                          (NOW(), NOW(), 'So cool!', 1, 4),
                                                                                          (NOW(), NOW(), 'Awesome Alice!', 1, 5),
                                                                                          (NOW(), NOW(), 'Nice writing!', 1, 2),
                                                                                          (NOW(), NOW(), 'Keep it up!', 1, 3),
                                                                                          (NOW(), NOW(), 'Love it!', 1, 4),
                                                                                          (NOW(), NOW(), 'Super helpful', 1, 5),
                                                                                          (NOW(), NOW(), 'Brilliant!', 1, 2),
                                                                                          (NOW(), NOW(), 'Learnt a lot', 1, 3),
                                                                                          (NOW(), NOW(), 'Fantastic read', 1, 4),
                                                                                          (NOW(), NOW(), '👏👏👏', 1, 5);

-- Post 2 (Alice) - 6
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Interesting!', 2, 3),
                                                                                          (NOW(), NOW(), 'Good job!', 2, 4),
                                                                                          (NOW(), NOW(), 'Very nice.', 2, 5),
                                                                                          (NOW(), NOW(), 'Cool!', 2, 2),
                                                                                          (NOW(), NOW(), 'Neat idea', 2, 3),
                                                                                          (NOW(), NOW(), 'Subtle point', 2, 4);

-- Post 3 (Alice) - 3
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Nice angle', 3, 2),
                                                                                          (NOW(), NOW(), 'Makes sense', 3, 5),
                                                                                          (NOW(), NOW(), 'Good read', 3, 4);

-- Post 5 (Alice) - 5
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Well written', 5, 2),
                                                                                          (NOW(), NOW(), 'Agree', 5, 3),
                                                                                          (NOW(), NOW(), 'Helpful', 5, 4),
                                                                                          (NOW(), NOW(), 'Beautiful', 5, 5),
                                                                                          (NOW(), NOW(), 'Thanks!', 5, 2);

-- Post 7 (Alice) - 2
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Short but sweet', 7, 3),
                                                                                          (NOW(), NOW(), 'Like it', 7, 5);

-- Post 8 (Bob) - 10
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Nice thoughts', 8, 1),
                                                                                          (NOW(), NOW(), 'Agree with you', 8, 3),
                                                                                          (NOW(), NOW(), 'Cool stuff', 8, 4),
                                                                                          (NOW(), NOW(), 'Insightful!', 8, 5),
                                                                                          (NOW(), NOW(), 'Well said!', 8, 1),
                                                                                          (NOW(), NOW(), 'Good point!', 8, 2),
                                                                                          (NOW(), NOW(), 'Smart take', 8, 3),
                                                                                          (NOW(), NOW(), 'Clear logic', 8, 4),
                                                                                          (NOW(), NOW(), '👌', 8, 5),
                                                                                          (NOW(), NOW(), 'True that', 8, 1);

-- Post 9 (Bob) - 3
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Solid', 9, 1),
                                                                                          (NOW(), NOW(), 'Nice read', 9, 4),
                                                                                          (NOW(), NOW(), 'Good!', 9, 5);

-- Post 11 (Bob) - 6
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Enjoyed this', 11, 1),
                                                                                          (NOW(), NOW(), 'Practical', 11, 3),
                                                                                          (NOW(), NOW(), 'Neatly done', 11, 4),
                                                                                          (NOW(), NOW(), 'Thanks Bob', 11, 5),
                                                                                          (NOW(), NOW(), 'Made my day', 11, 1),
                                                                                          (NOW(), NOW(), 'Solid work', 11, 2);

-- Post 12 (Bob) - 4
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Useful tip', 12, 1),
                                                                                          (NOW(), NOW(), 'Good summary', 12, 3),
                                                                                          (NOW(), NOW(), 'Like your view', 12, 4),
                                                                                          (NOW(), NOW(), 'Nice wrap-up', 12, 5);

-- Post 14 (Charlie) - 8
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Charlie rocks!', 14, 1),
                                                                                          (NOW(), NOW(), 'Deep thoughts', 14, 2),
                                                                                          (NOW(), NOW(), 'I agree', 14, 4),
                                                                                          (NOW(), NOW(), 'So true', 14, 5),
                                                                                          (NOW(), NOW(), 'Nice idea', 14, 1),
                                                                                          (NOW(), NOW(), 'Great!', 14, 2),
                                                                                          (NOW(), NOW(), 'Sharp point', 14, 4),
                                                                                          (NOW(), NOW(), 'Love this', 14, 5);

-- Post 15 (Charlie) - 4
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Concise', 15, 1),
                                                                                          (NOW(), NOW(), 'Clean', 15, 2),
                                                                                          (NOW(), NOW(), 'Valuable', 15, 4),
                                                                                          (NOW(), NOW(), '👏', 15, 5);

-- Post 16 (Charlie) - 2
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Good overview', 16, 1),
                                                                                          (NOW(), NOW(), 'Noted', 16, 2);

-- Post 19 (David) - 5
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Solid points', 19, 1),
                                                                                          (NOW(), NOW(), 'Agree', 19, 2),
                                                                                          (NOW(), NOW(), 'Thanks!', 19, 3),
                                                                                          (NOW(), NOW(), 'Nice one', 19, 5),
                                                                                          (NOW(), NOW(), 'Clear!', 19, 1);

-- Post 20 (David) - 4
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Helpful notes', 20, 1),
                                                                                          (NOW(), NOW(), 'Appreciated', 20, 2),
                                                                                          (NOW(), NOW(), 'Great tip', 20, 3),
                                                                                          (NOW(), NOW(), '🙌', 20, 5);

-- Post 22 (David) - 3
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Thanks David', 22, 1),
                                                                                          (NOW(), NOW(), 'That helps', 22, 3),
                                                                                          (NOW(), NOW(), 'Nice catch', 22, 5);

-- Post 23 (Eva) - 2
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
                                                                                          (NOW(), NOW(), 'Cool Eva', 23, 1),
                                                                                          (NOW(), NOW(), 'Good stuff', 23, 2);

-- Post 25 (Eva) - 1
INSERT INTO `comment` (created_at, updated_at, description, post_id, publish_user_id) VALUES
    (NOW(), NOW(), 'Nice 🙂', 25, 1);
