// AI Core Module for Learning Buddy Platform
// This module handles all AI-related functionality including prompt chaining,
// emotion detection, and buddy personality management

package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// BuddyPersonality represents different AI buddy personalities
type BuddyPersonality string

const (
	PersonalityMentor      BuddyPersonality = "mentor"
	PersonalityCheerleader BuddyPersonality = "cheerleader"
	PersonalityChill       BuddyPersonality = "chill"
	PersonalityFocused     BuddyPersonality = "focused"
)

// EmotionState represents the detected emotional state of the user
type EmotionState string

const (
	EmotionFrustrated EmotionState = "frustrated"
	EmotionMotivated  EmotionState = "motivated"
	EmotionConfused   EmotionState = "confused"
	EmotionConfident  EmotionState = "confident"
	EmotionTired      EmotionState = "tired"
	EmotionExcited    EmotionState = "excited"
)

// UserBehaviorData contains behavioral patterns for emotion detection
type UserBehaviorData struct {
	SessionDuration   int     `json:"session_duration"`   // minutes
	TaskFailures      int     `json:"task_failures"`      // number of failed attempts
	Retries           int     `json:"retries"`             // number of retries
	CompletionRate    float64 `json:"completion_rate"`     // percentage of completed tasks
	ResponseTime      int     `json:"response_time"`       // average response time in seconds
	StreakDays        int     `json:"streak_days"`         // current learning streak
	TimeOfDay         string  `json:"time_of_day"`         // morning, afternoon, evening, night
	LastActivity      string  `json:"last_activity"`       // timestamp of last activity
	RecentPerformance string  `json:"recent_performance"`  // improving, declining, stable
}

// PromptTemplate represents a template for AI prompts
type PromptTemplate struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Personality BuddyPersonality  `json:"personality"`
	Context     string            `json:"context"`
	Template    string            `json:"template"`
	Variables   map[string]string `json:"variables"`
	MaxTokens   int               `json:"max_tokens"`
	Temperature float64           `json:"temperature"`
}

// AIResponse represents the response from AI processing
type AIResponse struct {
	Message           string            `json:"message"`
	Personality       BuddyPersonality  `json:"personality"`
	DetectedEmotion   EmotionState      `json:"detected_emotion"`
	SuggestedActions  []string          `json:"suggested_actions"`
	XPReward          int               `json:"xp_reward"`
	ConfidenceScore   float64           `json:"confidence_score"`
	ProcessingTime    time.Duration     `json:"processing_time"`
	Metadata          map[string]string `json:"metadata"`
}

// PromptChainStep represents a step in the prompt chaining process
type PromptChainStep struct {
	StepName    string                 `json:"step_name"`
	Input       map[string]interface{} `json:"input"`
	Output      map[string]interface{} `json:"output"`
	Duration    time.Duration          `json:"duration"`
	Success     bool                   `json:"success"`
	ErrorMsg    string                 `json:"error_msg,omitempty"`
}

// PromptChainResult represents the result of a complete prompt chain execution
type PromptChainResult struct {
	ChainID     string            `json:"chain_id"`
	Steps       []PromptChainStep `json:"steps"`
	FinalResult AIResponse        `json:"final_result"`
	TotalTime   time.Duration     `json:"total_time"`
	Success     bool              `json:"success"`
}

// AICore is the main AI processing engine
type AICore struct {
	OpenAIKey       string
	PromptTemplates map[string]PromptTemplate
	PersonalityMap  map[BuddyPersonality]PersonalityConfig
}

// PersonalityConfig defines the configuration for each buddy personality
type PersonalityConfig struct {
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Traits          map[string]int    `json:"traits"`          // trait scores 1-10
	ResponseStyle   map[string]string `json:"response_style"`  // response characteristics
	PromptModifiers map[string]string `json:"prompt_modifiers"` // personality-specific prompt additions
	UnlockLevel     int               `json:"unlock_level"`
}

// NewAICore creates a new AI core instance
func NewAICore(openAIKey string) *AICore {
	core := &AICore{
		OpenAIKey:       openAIKey,
		PromptTemplates: make(map[string]PromptTemplate),
		PersonalityMap:  make(map[BuddyPersonality]PersonalityConfig),
	}
	
	core.initializePromptTemplates()
	core.initializePersonalities()
	
	return core
}

// initializePromptTemplates sets up the default prompt templates
func (ai *AICore) initializePromptTemplates() {
	templates := []PromptTemplate{
		{
			ID:          "emotion_detection",
			Name:        "Emotion Detection",
			Personality: PersonalityMentor,
			Context:     "Analyze user behavior to detect emotional state",
			Template: `Based on the following user behavior data, determine the user's emotional state:
Session Duration: {{.session_duration}} minutes
Task Failures: {{.task_failures}}
Retries: {{.retries}}
Completion Rate: {{.completion_rate}}%
Response Time: {{.response_time}} seconds
Streak: {{.streak_days}} days
Time of Day: {{.time_of_day}}
Recent Performance: {{.recent_performance}}

Analyze this data and respond with:
1. The most likely emotional state (frustrated, motivated, confused, confident, tired, excited)
2. Confidence level (0-100%)
3. Key behavioral indicators that led to this conclusion
4. Recommended buddy personality for this state

Format your response as JSON.`,
			Variables:   map[string]string{},
			MaxTokens:   300,
			Temperature: 0.3,
		},
		{
			ID:          "learning_suggestion",
			Name:        "Learning Suggestion Generator",
			Personality: PersonalityMentor,
			Context:     "Generate personalized learning suggestions",
			Template: `As a {{.personality}} learning buddy, create a personalized learning suggestion for a user with the following profile:

Current Level: {{.user_level}}
Learning Style: {{.learning_style}}
Preferred Difficulty: {{.preferred_difficulty}}/5
Recent Topics: {{.recent_topics}}
Emotional State: {{.emotional_state}}
Available Time: {{.available_time}} minutes

Create a learning activity that:
1. Matches their current skill level and preferences
2. Is appropriate for their emotional state
3. Can be completed in the available time
4. Includes clear steps and expected outcomes
5. Has built-in motivation and engagement elements

Respond in character as their {{.personality}} buddy with enthusiasm and encouragement.`,
			Variables:   map[string]string{},
			MaxTokens:   500,
			Temperature: 0.7,
		},
		{
			ID:          "progress_celebration",
			Name:        "Progress Celebration",
			Personality: PersonalityCheerleader,
			Context:     "Celebrate user achievements and progress",
			Template: `As an enthusiastic cheerleader buddy, celebrate the user's achievement:

Achievement: {{.achievement}}
XP Gained: {{.xp_gained}}
New Level: {{.new_level}}
Streak: {{.streak_days}} days
Recent Struggles: {{.recent_struggles}}

Create an encouraging and celebratory message that:
1. Acknowledges their specific achievement
2. Highlights their growth and progress
3. Motivates them to continue learning
4. Suggests what they might tackle next
5. Uses appropriate emojis and enthusiastic language

Keep the tone upbeat and genuinely excited about their success!`,
			Variables:   map[string]string{},
			MaxTokens:   400,
			Temperature: 0.8,
		},
		{
			ID:          "problem_solving_help",
			Name:        "Problem Solving Assistance",
			Personality: PersonalityFocused,
			Context:     "Help users solve complex problems step by step",
			Template: `As a focused, analytical buddy, help the user solve this problem:

Problem Description: {{.problem_description}}
User's Attempted Solutions: {{.attempted_solutions}}
Error Messages: {{.error_messages}}
User's Experience Level: {{.experience_level}}
Available Resources: {{.available_resources}}

Provide systematic help by:
1. Breaking down the problem into smaller components
2. Identifying the root cause based on their attempts
3. Suggesting a step-by-step solution approach
4. Explaining the reasoning behind each step
5. Anticipating potential pitfalls and how to avoid them

Be precise, logical, and thorough in your analysis.`,
			Variables:   map[string]string{},
			MaxTokens:   600,
			Temperature: 0.4,
		},
		{
			ID:          "motivation_boost",
			Name:        "Motivation and Encouragement",
			Personality: PersonalityChill,
			Context:     "Provide gentle motivation and stress relief",
			Template: `As a chill, supportive friend, help the user who is feeling {{.emotional_state}}:

Current Situation: {{.current_situation}}
Recent Challenges: {{.recent_challenges}}
User's Goals: {{.user_goals}}
Stress Level: {{.stress_level}}/10
Time Pressure: {{.time_pressure}}

Provide gentle support by:
1. Acknowledging their feelings without judgment
2. Offering a fresh perspective on their situation
3. Suggesting small, manageable next steps
4. Reminding them of their past successes
5. Encouraging self-care and balance

Keep the tone relaxed, understanding, and genuinely caring.`,
			Variables:   map[string]string{},
			MaxTokens:   450,
			Temperature: 0.6,
		},
	}
	
	for _, template := range templates {
		ai.PromptTemplates[template.ID] = template
	}
}

// initializePersonalities sets up the personality configurations
func (ai *AICore) initializePersonalities() {
	personalities := map[BuddyPersonality]PersonalityConfig{
		PersonalityMentor: {
			Name:        "Wise Mentor",
			Description: "Patient teacher who guides you through challenges with wisdom and experience",
			Traits: map[string]int{
				"patience":      9,
				"wisdom":        10,
				"encouragement": 7,
				"directness":    8,
				"empathy":       8,
			},
			ResponseStyle: map[string]string{
				"teaching":   "step-by-step guidance",
				"feedback":   "constructive and detailed",
				"motivation": "growth-focused",
				"tone":       "wise and patient",
			},
			PromptModifiers: map[string]string{
				"prefix": "As a wise and experienced mentor,",
				"style":  "Provide thoughtful, step-by-step guidance with clear explanations.",
				"suffix": "Remember to be patient and encouraging in your response.",
			},
			UnlockLevel: 1,
		},
		PersonalityCheerleader: {
			Name:        "Enthusiastic Cheerleader",
			Description: "Energetic supporter who celebrates every victory and keeps you motivated",
			Traits: map[string]int{
				"enthusiasm": 10,
				"positivity": 10,
				"energy":     9,
				"empathy":    8,
				"optimism":   10,
			},
			ResponseStyle: map[string]string{
				"teaching":   "encouraging and uplifting",
				"feedback":   "positive and motivating",
				"motivation": "celebration-focused",
				"tone":       "enthusiastic and energetic",
			},
			PromptModifiers: map[string]string{
				"prefix": "As an enthusiastic and supportive cheerleader,",
				"style":  "Use encouraging language, celebrate achievements, and maintain high energy.",
				"suffix": "End with motivation and excitement about their progress!",
			},
			UnlockLevel: 1,
		},
		PersonalityChill: {
			Name:        "Chill Friend",
			Description: "Relaxed companion who keeps things light, fun, and stress-free",
			Traits: map[string]int{
				"calmness":    9,
				"humor":       8,
				"flexibility": 10,
				"casualness":  9,
				"empathy":     9,
			},
			ResponseStyle: map[string]string{
				"teaching":   "casual and relaxed",
				"feedback":   "gentle and non-judgmental",
				"motivation": "fun-focused and low-pressure",
				"tone":       "laid-back and friendly",
			},
			PromptModifiers: map[string]string{
				"prefix": "As a chill, laid-back friend,",
				"style":  "Keep things casual, use humor when appropriate, and reduce pressure.",
				"suffix": "Remember to keep it light and stress-free!",
			},
			UnlockLevel: 3,
		},
		PersonalityFocused: {
			Name:        "Focused Analyst",
			Description: "Detail-oriented problem solver who excels at breaking down complex challenges",
			Traits: map[string]int{
				"precision":    10,
				"logic":        10,
				"focus":        9,
				"thoroughness": 9,
				"analytical":   10,
			},
			ResponseStyle: map[string]string{
				"teaching":   "systematic and analytical",
				"feedback":   "detailed and precise",
				"motivation": "problem-solving focused",
				"tone":       "focused and methodical",
			},
			PromptModifiers: map[string]string{
				"prefix": "As a focused, analytical problem-solver,",
				"style":  "Be systematic, precise, and thorough in your analysis.",
				"suffix": "Ensure your response is logically structured and comprehensive.",
			},
			UnlockLevel: 5,
		},
	}
	
	ai.PersonalityMap = personalities
}

// DetectEmotion analyzes user behavior to detect emotional state
func (ai *AICore) DetectEmotion(ctx context.Context, behaviorData UserBehaviorData) (EmotionState, float64, error) {
	// Simple rule-based emotion detection (in production, use ML model or LLM)
	
	// Calculate frustration indicators
	frustrationScore := 0.0
	if behaviorData.TaskFailures > 3 {
		frustrationScore += 0.3
	}
	if behaviorData.Retries > 5 {
		frustrationScore += 0.2
	}
	if behaviorData.CompletionRate < 0.5 {
		frustrationScore += 0.3
	}
	if behaviorData.ResponseTime > 30 {
		frustrationScore += 0.2
	}
	
	// Calculate motivation indicators
	motivationScore := 0.0
	if behaviorData.StreakDays > 3 {
		motivationScore += 0.3
	}
	if behaviorData.CompletionRate > 0.8 {
		motivationScore += 0.3
	}
	if behaviorData.SessionDuration > 30 {
		motivationScore += 0.2
	}
	if behaviorData.RecentPerformance == "improving" {
		motivationScore += 0.2
	}
	
	// Calculate confusion indicators
	confusionScore := 0.0
	if behaviorData.TaskFailures > 1 && behaviorData.Retries > 3 {
		confusionScore += 0.4
	}
	if behaviorData.ResponseTime > 60 {
		confusionScore += 0.3
	}
	if behaviorData.CompletionRate < 0.3 {
		confusionScore += 0.3
	}
	
	// Calculate fatigue indicators
	fatigueScore := 0.0
	if behaviorData.SessionDuration > 120 {
		fatigueScore += 0.3
	}
	if behaviorData.TimeOfDay == "night" {
		fatigueScore += 0.2
	}
	if behaviorData.RecentPerformance == "declining" {
		fatigueScore += 0.3
	}
	if behaviorData.ResponseTime > 45 {
		fatigueScore += 0.2
	}
	
	// Determine dominant emotion
	scores := map[EmotionState]float64{
		EmotionFrustrated: frustrationScore,
		EmotionMotivated:  motivationScore,
		EmotionConfused:   confusionScore,
		EmotionTired:      fatigueScore,
	}
	
	// Find highest scoring emotion
	var dominantEmotion EmotionState
	var highestScore float64
	
	for emotion, score := range scores {
		if score > highestScore {
			highestScore = score
			dominantEmotion = emotion
		}
	}
	
	// Default to confident if no strong indicators
	if highestScore < 0.4 {
		dominantEmotion = EmotionConfident
		highestScore = 0.6
	}
	
	return dominantEmotion, highestScore, nil
}

// ExecutePromptChain runs a complete prompt chain for AI processing
func (ai *AICore) ExecutePromptChain(ctx context.Context, chainType string, input map[string]interface{}) (*PromptChainResult, error) {
	startTime := time.Now()
	chainID := fmt.Sprintf("%s_%d", chainType, startTime.Unix())
	
	result := &PromptChainResult{
		ChainID: chainID,
		Steps:   []PromptChainStep{},
		Success: true,
	}
	
	switch chainType {
	case "detect_adapt_suggest_reward":
		return ai.executeDetectAdaptSuggestReward(ctx, input, result)
	case "problem_solving":
		return ai.executeProblemSolvingChain(ctx, input, result)
	case "motivation_boost":
		return ai.executeMotivationBoostChain(ctx, input, result)
	default:
		return nil, fmt.Errorf("unknown chain type: %s", chainType)
	}
}

// executeDetectAdaptSuggestReward implements the core prompt chain: detect → adapt → suggest → reward
func (ai *AICore) executeDetectAdaptSuggestReward(ctx context.Context, input map[string]interface{}, result *PromptChainResult) (*PromptChainResult, error) {
	// Step 1: Detect emotion
	stepStart := time.Now()
	behaviorData := input["behavior_data"].(UserBehaviorData)
	emotion, confidence, err := ai.DetectEmotion(ctx, behaviorData)
	
	step1 := PromptChainStep{
		StepName: "detect_emotion",
		Input:    map[string]interface{}{"behavior_data": behaviorData},
		Output:   map[string]interface{}{"emotion": emotion, "confidence": confidence},
		Duration: time.Since(stepStart),
		Success:  err == nil,
	}
	if err != nil {
		step1.ErrorMsg = err.Error()
		result.Success = false
	}
	result.Steps = append(result.Steps, step1)
	
	// Step 2: Adapt personality
	stepStart = time.Now()
	personality := ai.adaptPersonalityToEmotion(emotion)
	
	step2 := PromptChainStep{
		StepName: "adapt_personality",
		Input:    map[string]interface{}{"emotion": emotion},
		Output:   map[string]interface{}{"personality": personality},
		Duration: time.Since(stepStart),
		Success:  true,
	}
	result.Steps = append(result.Steps, step2)
	
	// Step 3: Generate suggestion
	stepStart = time.Now()
	suggestion, err := ai.generateLearningResponse(ctx, personality, emotion, input)
	
	step3 := PromptChainStep{
		StepName: "generate_suggestion",
		Input:    map[string]interface{}{"personality": personality, "emotion": emotion, "context": input},
		Output:   map[string]interface{}{"suggestion": suggestion},
		Duration: time.Since(stepStart),
		Success:  err == nil,
	}
	if err != nil {
		step3.ErrorMsg = err.Error()
		result.Success = false
	}
	result.Steps = append(result.Steps, step3)
	
	// Step 4: Calculate reward
	stepStart = time.Now()
	xpReward := ai.calculateXPReward(emotion, confidence, len(result.Steps))
	
	step4 := PromptChainStep{
		StepName: "calculate_reward",
		Input:    map[string]interface{}{"emotion": emotion, "confidence": confidence},
		Output:   map[string]interface{}{"xp_reward": xpReward},
		Duration: time.Since(stepStart),
		Success:  true,
	}
	result.Steps = append(result.Steps, step4)
	
	// Compile final result
	result.FinalResult = AIResponse{
		Message:         suggestion,
		Personality:     personality,
		DetectedEmotion: emotion,
		XPReward:        xpReward,
		ConfidenceScore: confidence,
		ProcessingTime:  time.Since(time.Now().Add(-result.TotalTime)),
	}
	
	result.TotalTime = time.Since(time.Now().Add(-result.TotalTime))
	return result, nil
}

// Helper methods for prompt chain execution

func (ai *AICore) adaptPersonalityToEmotion(emotion EmotionState) BuddyPersonality {
	switch emotion {
	case EmotionFrustrated, EmotionConfused:
		return PersonalityMentor
	case EmotionTired:
		return PersonalityChill
	case EmotionMotivated, EmotionExcited:
		return PersonalityCheerleader
	case EmotionConfident:
		return PersonalityFocused
	default:
		return PersonalityMentor
	}
}

func (ai *AICore) generateLearningResponse(ctx context.Context, personality BuddyPersonality, emotion EmotionState, input map[string]interface{}) (string, error) {
	// In a real implementation, this would call OpenAI API
	// For now, return a mock response based on personality and emotion
	
	responses := map[BuddyPersonality]map[EmotionState]string{
		PersonalityMentor: {
			EmotionFrustrated: "I can see you're facing some challenges. Let's break this down step by step and work through it together.",
			EmotionConfused:   "It's completely normal to feel confused when learning something new. Let me help clarify the concepts for you.",
			EmotionTired:      "You've been working hard! Let's take a gentler approach and focus on one small step at a time.",
		},
		PersonalityCheerleader: {
			EmotionMotivated: "You're absolutely crushing it! Your motivation is inspiring! Let's channel that energy into your next challenge! 🎉",
			EmotionExcited:   "I love your enthusiasm! You're ready to take on anything! Let's dive into something exciting! ⭐",
		},
		PersonalityChill: {
			EmotionTired:      "Hey, no pressure at all. Let's just take it easy and explore something fun and relaxed.",
			EmotionFrustrated: "All good, we've all been there. Let's step back and approach this from a different angle.",
		},
		PersonalityFocused: {
			EmotionConfident: "Great! You're in the zone. Let's tackle something challenging that will really test your skills.",
		},
	}
	
	if personalityResponses, exists := responses[personality]; exists {
		if response, exists := personalityResponses[emotion]; exists {
			return response, nil
		}
	}
	
	// Default response
	return "I'm here to help you learn and grow. What would you like to work on today?", nil
}

func (ai *AICore) calculateXPReward(emotion EmotionState, confidence float64, chainSteps int) int {
	baseXP := 25
	
	// Emotion-based multiplier
	emotionMultiplier := map[EmotionState]float64{
		EmotionFrustrated: 1.5, // Extra reward for perseverance
		EmotionConfused:   1.3, // Reward for seeking help
		EmotionMotivated:  1.2, // Standard motivated reward
		EmotionConfident:  1.0, // Standard reward
		EmotionTired:      1.4, // Reward for continuing despite fatigue
		EmotionExcited:    1.1, // Small bonus for enthusiasm
	}
	
	multiplier := emotionMultiplier[emotion]
	if multiplier == 0 {
		multiplier = 1.0
	}
	
	// Confidence bonus
	confidenceBonus := confidence * 0.5
	
	// Chain complexity bonus
	complexityBonus := float64(chainSteps) * 0.1
	
	finalXP := float64(baseXP) * multiplier * (1.0 + confidenceBonus + complexityBonus)
	return int(finalXP)
}

// Additional helper methods would be implemented here for:
// - executeProblemSolvingChain
// - executeMotivationBoostChain
// - OpenAI API integration
// - Template variable substitution
// - Response validation and filtering

