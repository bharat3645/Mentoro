-- Learning Buddy Platform Database Schema
-- PostgreSQL Database Initialization Script

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    uuid UUID DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    level INTEGER DEFAULT 1,
    total_xp INTEGER DEFAULT 0,
    current_streak INTEGER DEFAULT 0,
    max_streak INTEGER DEFAULT 0,
    last_activity TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    buddy_mood VARCHAR(20) DEFAULT 'mentor',
    preferences JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Quests table
CREATE TABLE quests (
    id SERIAL PRIMARY KEY,
    uuid UUID DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    quest_type VARCHAR(50) NOT NULL, -- 'code', 'focus', 'learn', 'debug', 'test'
    difficulty INTEGER DEFAULT 1, -- 1-5 scale
    total_tasks INTEGER DEFAULT 1,
    completed_tasks INTEGER DEFAULT 0,
    xp_reward INTEGER DEFAULT 50,
    status VARCHAR(20) DEFAULT 'active', -- 'active', 'completed', 'paused', 'failed'
    due_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Quest tasks table (sub-tasks within a quest)
CREATE TABLE quest_tasks (
    id SERIAL PRIMARY KEY,
    quest_id INTEGER REFERENCES quests(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    order_index INTEGER DEFAULT 0,
    is_completed BOOLEAN DEFAULT FALSE,
    completed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Badges table
CREATE TABLE badges (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    icon VARCHAR(10), -- emoji or icon identifier
    category VARCHAR(50), -- 'achievement', 'streak', 'skill', 'social'
    requirements JSONB, -- JSON object defining requirements
    xp_reward INTEGER DEFAULT 0,
    rarity VARCHAR(20) DEFAULT 'common', -- 'common', 'rare', 'epic', 'legendary'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- User badges (earned badges)
CREATE TABLE user_badges (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    badge_id INTEGER REFERENCES badges(id) ON DELETE CASCADE,
    earned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, badge_id)
);

-- XP transactions (track all XP gains/losses)
CREATE TABLE xp_transactions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    amount INTEGER NOT NULL, -- can be negative for penalties
    source_type VARCHAR(50) NOT NULL, -- 'quest', 'badge', 'streak', 'penalty'
    source_id INTEGER, -- ID of the quest, badge, etc.
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Buddy conversations
CREATE TABLE buddy_conversations (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    user_message TEXT NOT NULL,
    buddy_response TEXT NOT NULL,
    buddy_mood VARCHAR(20) NOT NULL,
    context JSONB, -- Additional context for the conversation
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- User sessions (track activity patterns)
CREATE TABLE user_sessions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    session_start TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    session_end TIMESTAMP,
    duration_minutes INTEGER,
    tasks_completed INTEGER DEFAULT 0,
    tasks_failed INTEGER DEFAULT 0,
    retries INTEGER DEFAULT 0,
    mood_detected VARCHAR(20),
    xp_gained INTEGER DEFAULT 0
);

-- Learning paths (structured learning sequences)
CREATE TABLE learning_paths (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    difficulty_level INTEGER DEFAULT 1,
    estimated_hours INTEGER,
    prerequisites JSONB,
    tags JSONB,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- User progress in learning paths
CREATE TABLE user_learning_progress (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    learning_path_id INTEGER REFERENCES learning_paths(id) ON DELETE CASCADE,
    progress_percentage INTEGER DEFAULT 0,
    current_step INTEGER DEFAULT 1,
    started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP,
    UNIQUE(user_id, learning_path_id)
);

-- DIY (Do It Yourself) projects
CREATE TABLE diy_projects (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    project_type VARCHAR(50), -- 'coding', 'design', 'research', 'writing'
    difficulty INTEGER DEFAULT 1,
    estimated_hours INTEGER,
    status VARCHAR(20) DEFAULT 'planning', -- 'planning', 'in_progress', 'completed', 'abandoned'
    resources JSONB, -- Links, files, references
    milestones JSONB, -- Project milestones
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP
);

-- Feedback and ratings
CREATE TABLE feedback (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    feedback_type VARCHAR(50) NOT NULL, -- 'quest', 'buddy', 'feature', 'bug'
    target_id INTEGER, -- ID of quest, conversation, etc.
    rating INTEGER CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_level ON users(level);
CREATE INDEX idx_quests_user_id ON quests(user_id);
CREATE INDEX idx_quests_status ON quests(status);
CREATE INDEX idx_quest_tasks_quest_id ON quest_tasks(quest_id);
CREATE INDEX idx_user_badges_user_id ON user_badges(user_id);
CREATE INDEX idx_xp_transactions_user_id ON xp_transactions(user_id);
CREATE INDEX idx_buddy_conversations_user_id ON buddy_conversations(user_id);
CREATE INDEX idx_user_sessions_user_id ON user_sessions(user_id);
CREATE INDEX idx_user_learning_progress_user_id ON user_learning_progress(user_id);
CREATE INDEX idx_diy_projects_user_id ON diy_projects(user_id);
CREATE INDEX idx_feedback_user_id ON feedback(user_id);

-- Create triggers for updated_at timestamps
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_quests_updated_at BEFORE UPDATE ON quests
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Insert sample badges
INSERT INTO badges (name, description, icon, category, requirements, xp_reward, rarity) VALUES
('Code Warrior', 'Fixed 10 bugs in your code', '⚔️', 'achievement', '{"bugs_fixed": 10}', 100, 'common'),
('Focus Master', 'Completed 50 focus sessions', '🎯', 'achievement', '{"focus_sessions": 50}', 200, 'rare'),
('Streak Champion', 'Maintained a 30-day learning streak', '🔥', 'streak', '{"max_streak": 30}', 500, 'epic'),
('Bug Hunter', 'Found and reported 5 bugs', '🐛', 'achievement', '{"bugs_reported": 5}', 75, 'common'),
('Level Achiever', 'Reached level 10', '⭐', 'achievement', '{"level": 10}', 300, 'rare'),
('XP Collector', 'Earned 5000 total XP', '💎', 'achievement', '{"total_xp": 5000}', 250, 'rare'),
('Quest Completer', 'Completed 25 quests', '🏆', 'achievement', '{"quests_completed": 25}', 400, 'epic'),
('Learning Machine', 'Learned 15 new concepts', '🧠', 'skill', '{"concepts_learned": 15}', 350, 'rare'),
('Early Bird', 'Started learning before 8 AM', '🌅', 'achievement', '{"early_sessions": 10}', 50, 'common'),
('Night Owl', 'Learned after 10 PM', '🦉', 'achievement', '{"night_sessions": 10}', 50, 'common'),
('Social Learner', 'Helped 5 other learners', '🤝', 'social', '{"helped_users": 5}', 150, 'common'),
('Perfectionist', 'Completed 10 quests with 100% accuracy', '💯', 'achievement', '{"perfect_quests": 10}', 300, 'epic');

-- Insert sample learning paths
INSERT INTO learning_paths (title, description, difficulty_level, estimated_hours, prerequisites, tags) VALUES
('JavaScript Fundamentals', 'Learn the basics of JavaScript programming', 1, 20, '[]', '["javascript", "programming", "beginner"]'),
('React Development', 'Build modern web applications with React', 2, 30, '["JavaScript Fundamentals"]', '["react", "frontend", "javascript"]'),
('Backend with Node.js', 'Create server-side applications with Node.js', 2, 25, '["JavaScript Fundamentals"]', '["nodejs", "backend", "javascript"]'),
('Database Design', 'Learn to design and work with databases', 2, 15, '[]', '["database", "sql", "design"]'),
('Full-Stack Development', 'Complete full-stack web development', 3, 50, '["React Development", "Backend with Node.js", "Database Design"]', '["fullstack", "web", "advanced"]');

-- Insert a sample user for testing
INSERT INTO users (username, email, password_hash, first_name, last_name, level, total_xp, current_streak, buddy_mood) VALUES
('alex_learner', 'alex@example.com', '$2a$10$example_hash', 'Alex', 'Smith', 5, 750, 7, 'focused');

-- Insert sample quests for the test user
INSERT INTO quests (user_id, title, description, quest_type, difficulty, total_tasks, completed_tasks, xp_reward, status) VALUES
(1, 'Fix 3 bugs', 'Debug and fix 3 code issues in your project', 'code', 2, 3, 2, 150, 'active'),
(1, '10 min focus session', 'Complete a focused work session without distractions', 'focus', 1, 10, 7, 50, 'active'),
(1, 'Learn new concept', 'Study and understand a new programming concept', 'learn', 3, 1, 0, 200, 'active');

-- Insert sample quest tasks
INSERT INTO quest_tasks (quest_id, title, description, order_index, is_completed) VALUES
(1, 'Identify first bug', 'Find the first bug in the codebase', 1, true),
(1, 'Fix first bug', 'Implement a solution for the first bug', 2, true),
(1, 'Identify second bug', 'Find the second bug in the codebase', 3, false),
(2, 'Set up workspace', 'Prepare your workspace for focused work', 1, true),
(2, 'Start focus timer', 'Begin the 10-minute focus session', 2, true),
(2, 'Complete focus work', 'Work without distractions for the full session', 3, false);

-- Insert sample XP transactions
INSERT INTO xp_transactions (user_id, amount, source_type, source_id, description) VALUES
(1, 100, 'quest', 1, 'Completed quest task: Identify first bug'),
(1, 50, 'quest', 2, 'Completed quest task: Set up workspace'),
(1, 25, 'streak', NULL, 'Daily streak bonus'),
(1, 75, 'badge', 1, 'Earned Code Warrior badge');

-- Insert sample buddy conversations
INSERT INTO buddy_conversations (user_id, user_message, buddy_response, buddy_mood, context) VALUES
(1, 'I''m feeling stuck on this problem', 'I understand! Let''s break it down into smaller steps. What specific part is challenging you?', 'mentor', '{"session_duration": 15, "task_failures": 1}'),
(1, 'Great job today!', 'You''re absolutely crushing it! This question shows you''re really engaged! 🎉', 'cheerleader', '{"session_duration": 45, "tasks_completed": 3}');

-- Grant necessary permissions (adjust as needed for your setup)
-- GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO buddy_user;
-- GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO buddy_user;

