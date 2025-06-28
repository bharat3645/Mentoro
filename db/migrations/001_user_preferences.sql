-- Migration: Add user preferences and settings
-- Version: 001
-- Date: 2024-01-01

-- Add new columns to users table for enhanced personalization
ALTER TABLE users ADD COLUMN IF NOT EXISTS timezone VARCHAR(50) DEFAULT 'UTC';
ALTER TABLE users ADD COLUMN IF NOT EXISTS notification_preferences JSONB DEFAULT '{"email": true, "push": true, "buddy_reminders": true}';
ALTER TABLE users ADD COLUMN IF NOT EXISTS learning_style VARCHAR(50) DEFAULT 'balanced'; -- 'visual', 'auditory', 'kinesthetic', 'balanced'
ALTER TABLE users ADD COLUMN IF NOT EXISTS daily_goal_minutes INTEGER DEFAULT 30;
ALTER TABLE users ADD COLUMN IF NOT EXISTS preferred_difficulty INTEGER DEFAULT 2; -- 1-5 scale

-- Add buddy personality settings
CREATE TABLE IF NOT EXISTS buddy_personalities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    traits JSONB, -- personality traits and characteristics
    response_style JSONB, -- how the buddy responds in different situations
    unlock_level INTEGER DEFAULT 1, -- level required to unlock this personality
    is_default BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert default buddy personalities
INSERT INTO buddy_personalities (name, description, traits, response_style, unlock_level, is_default) VALUES
('Mentor', 'Wise and patient teacher who guides you through challenges', 
 '{"patience": 9, "wisdom": 10, "encouragement": 7, "directness": 8}',
 '{"teaching": "step-by-step", "feedback": "constructive", "motivation": "growth-focused"}', 1, true),
('Cheerleader', 'Enthusiastic supporter who celebrates every victory', 
 '{"enthusiasm": 10, "positivity": 10, "energy": 9, "empathy": 8}',
 '{"teaching": "encouraging", "feedback": "positive", "motivation": "celebration-focused"}', 1, false),
('Chill Friend', 'Relaxed companion who keeps things light and fun', 
 '{"calmness": 9, "humor": 8, "flexibility": 10, "casualness": 9}',
 '{"teaching": "casual", "feedback": "gentle", "motivation": "fun-focused"}', 3, false),
('Focused Analyst', 'Detail-oriented helper who loves solving complex problems', 
 '{"precision": 10, "logic": 10, "focus": 9, "thoroughness": 9}',
 '{"teaching": "analytical", "feedback": "detailed", "motivation": "problem-solving"}', 5, false);

-- Add user's selected buddy personality
ALTER TABLE users ADD COLUMN IF NOT EXISTS buddy_personality_id INTEGER REFERENCES buddy_personalities(id) DEFAULT 1;

-- Create table for tracking mood detection patterns
CREATE TABLE IF NOT EXISTS mood_detection_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    session_id INTEGER REFERENCES user_sessions(id) ON DELETE CASCADE,
    detected_mood VARCHAR(20),
    confidence_score DECIMAL(3,2), -- 0.00 to 1.00
    factors JSONB, -- factors that influenced the detection
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for new tables
CREATE INDEX IF NOT EXISTS idx_buddy_personalities_unlock_level ON buddy_personalities(unlock_level);
CREATE INDEX IF NOT EXISTS idx_users_buddy_personality ON users(buddy_personality_id);
CREATE INDEX IF NOT EXISTS idx_mood_detection_user_session ON mood_detection_logs(user_id, session_id);

-- Update existing sample user with new fields
UPDATE users SET 
    timezone = 'America/New_York',
    learning_style = 'balanced',
    daily_goal_minutes = 45,
    preferred_difficulty = 3,
    buddy_personality_id = 1
WHERE username = 'alex_learner';

