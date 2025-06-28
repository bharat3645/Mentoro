// OpenAI Integration for Learning Buddy Platform
// This module provides integration with OpenAI API for AI buddy functionality

package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// OpenAIClient handles communication with OpenAI API
type OpenAIClient struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

// OpenAIRequest represents a request to OpenAI API
type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
	TopP        float64   `json:"top_p,omitempty"`
	Stream      bool      `json:"stream,omitempty"`
}

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`    // "system", "user", "assistant"
	Content string `json:"content"`
}

// OpenAIResponse represents the response from OpenAI API
type OpenAIResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// Choice represents a response choice
type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

// Usage represents token usage information
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// NewOpenAIClient creates a new OpenAI client
func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{
		APIKey:  apiKey,
		BaseURL: "https://api.openai.com/v1",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GenerateResponse sends a request to OpenAI and returns the response
func (client *OpenAIClient) GenerateResponse(ctx context.Context, req OpenAIRequest) (*OpenAIResponse, error) {
	// Set default model if not specified
	if req.Model == "" {
		req.Model = "gpt-4"
	}

	// Prepare request body
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, "POST", client.BaseURL+"/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+client.APIKey)

	// Send request
	resp, err := client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var openAIResp OpenAIResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &openAIResp, nil
}

// Enhanced AI Core with OpenAI Integration

// GenerateResponseWithOpenAI generates an AI response using OpenAI API
func (ai *AICore) GenerateResponseWithOpenAI(ctx context.Context, templateID string, variables map[string]interface{}, personality BuddyPersonality) (*AIResponse, error) {
	// Get template
	template, exists := ai.PromptTemplates[templateID]
	if !exists {
		return nil, fmt.Errorf("template not found: %s", templateID)
	}

	// Get personality config
	personalityConfig, exists := ai.PersonalityMap[personality]
	if !exists {
		return nil, fmt.Errorf("personality not found: %s", personality)
	}

	// Substitute variables in template
	prompt := ai.substituteVariables(template.Template, variables)

	// Add personality modifiers
	systemPrompt := fmt.Sprintf("%s %s %s", 
		personalityConfig.PromptModifiers["prefix"],
		personalityConfig.PromptModifiers["style"],
		personalityConfig.PromptModifiers["suffix"])

	// Create OpenAI client (in production, this would be initialized once)
	client := NewOpenAIClient(ai.OpenAIKey)

	// Prepare OpenAI request
	openAIReq := OpenAIRequest{
		Model: "gpt-4",
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: prompt},
		},
		MaxTokens:   template.MaxTokens,
		Temperature: template.Temperature,
	}

	// Generate response
	startTime := time.Now()
	openAIResp, err := client.GenerateResponse(ctx, openAIReq)
	if err != nil {
		return nil, fmt.Errorf("OpenAI request failed: %w", err)
	}
	processingTime := time.Since(startTime)

	// Extract response content
	if len(openAIResp.Choices) == 0 {
		return nil, fmt.Errorf("no response choices returned")
	}

	responseContent := openAIResp.Choices[0].Message.Content

	// Create AI response
	aiResponse := &AIResponse{
		Message:         responseContent,
		Personality:     personality,
		ProcessingTime:  processingTime,
		ConfidenceScore: 0.85, // Default confidence, could be calculated based on response quality
		Metadata: map[string]string{
			"template_id":       templateID,
			"model":            openAIResp.Model,
			"total_tokens":     fmt.Sprintf("%d", openAIResp.Usage.TotalTokens),
			"prompt_tokens":    fmt.Sprintf("%d", openAIResp.Usage.PromptTokens),
			"completion_tokens": fmt.Sprintf("%d", openAIResp.Usage.CompletionTokens),
		},
	}

	return aiResponse, nil
}

// substituteVariables replaces template variables with actual values
func (ai *AICore) substituteVariables(template string, variables map[string]interface{}) string {
	result := template
	
	for key, value := range variables {
		placeholder := fmt.Sprintf("{{.%s}}", key)
		replacement := fmt.Sprintf("%v", value)
		result = strings.ReplaceAll(result, placeholder, replacement)
	}
	
	return result
}

// Mock AI Response Generator (for development/testing without OpenAI API)
func (ai *AICore) GenerateMockResponse(ctx context.Context, templateID string, variables map[string]interface{}, personality BuddyPersonality) (*AIResponse, error) {
	// Mock responses based on template and personality
	mockResponses := map[string]map[BuddyPersonality]string{
		"emotion_detection": {
			PersonalityMentor: `{
				"emotional_state": "focused",
				"confidence": 0.78,
				"indicators": ["consistent session duration", "low failure rate", "steady progress"],
				"recommended_personality": "mentor",
				"reasoning": "User shows signs of focused learning with good progress patterns"
			}`,
		},
		"learning_suggestion_mentor": {
			PersonalityMentor: "I can see you're ready to tackle something challenging! Based on your current progress, I recommend we work on debugging techniques. Let's start with a systematic approach: first, reproduce the issue consistently, then isolate the problem area, and finally implement a targeted fix. This will build your problem-solving skills step by step.",
		},
		"learning_suggestion_cheerleader": {
			PersonalityCheerleader: "🎉 You're absolutely crushing it! Your progress has been amazing! Let's channel that awesome energy into a fun coding challenge! I've got the perfect project that will let you show off your skills while learning something new. You're going to love this - it combines everything you've been working on! Ready to create something incredible? 💪⭐",
		},
		"learning_suggestion_chill": {
			PersonalityChill: "Hey there! No pressure at all, but I found something pretty cool we could explore together. It's this relaxed coding exercise that's actually quite fun - think of it like a puzzle rather than work. We can take our time, maybe grab a coffee, and just see where it takes us. Sound good?",
		},
		"learning_suggestion_focused": {
			PersonalityFocused: "Let's dive deep into this systematically. Based on your current skill level and learning objectives, I recommend a structured approach to advanced algorithms. We'll start with complexity analysis, move through optimization techniques, and conclude with practical implementation. Each phase builds on the previous one, ensuring comprehensive understanding.",
		},
		"problem_solving_help": {
			PersonalityFocused: "Let me help you break this down systematically. First, let's identify the core issue: it appears to be a logic error in your conditional statements. Here's my analysis: 1) The condition is checking the wrong variable, 2) The loop termination isn't properly handled, 3) Edge cases aren't covered. Let's fix these one by one with a methodical approach.",
		},
		"progress_celebration": {
			PersonalityCheerleader: "🎉 INCREDIBLE WORK! You just leveled up and I am SO proud of you! 🌟 Look at how far you've come - from struggling with basic concepts to solving complex problems like a pro! Your dedication is absolutely inspiring! This achievement shows you're ready for even bigger challenges. What amazing thing should we tackle next? I'm so excited to see what you'll accomplish! 💪✨",
		},
		"motivation_boost": {
			PersonalityChill: "Hey, I totally get it - we all have those tough days where everything feels overwhelming. 💙 But here's the thing: you've made it through 100% of your difficult days so far, and that's pretty amazing! Remember when you thought that last challenge was impossible? Look at you now! Let's just focus on one tiny step today. No pressure, just progress. You've got this, and I've got your back. 🌱",
		},
	}

	// Get mock response
	templateResponses, exists := mockResponses[templateID]
	if !exists {
		templateResponses = mockResponses["learning_suggestion_mentor"] // Default fallback
	}

	response, exists := templateResponses[personality]
	if !exists {
		// Fallback to any available response for this template
		for _, resp := range templateResponses {
			response = resp
			break
		}
	}

	// If still no response, provide a generic one
	if response == "" {
		response = "I'm here to help you learn and grow! What would you like to work on today?"
	}

	// Create AI response
	aiResponse := &AIResponse{
		Message:         response,
		Personality:     personality,
		ProcessingTime:  time.Millisecond * 150, // Simulate processing time
		ConfidenceScore: 0.75,
		Metadata: map[string]string{
			"template_id": templateID,
			"mode":       "mock",
			"personality": string(personality),
		},
	}

	// Add some personality-specific metadata
	switch personality {
	case PersonalityCheerleader:
		aiResponse.XPReward = 25
		aiResponse.SuggestedActions = []string{"celebrate", "share_progress", "set_new_goal"}
	case PersonalityMentor:
		aiResponse.XPReward = 30
		aiResponse.SuggestedActions = []string{"practice", "review_concepts", "ask_questions"}
	case PersonalityChill:
		aiResponse.XPReward = 20
		aiResponse.SuggestedActions = []string{"take_break", "reflect", "gentle_practice"}
	case PersonalityFocused:
		aiResponse.XPReward = 35
		aiResponse.SuggestedActions = []string{"deep_dive", "analyze", "optimize"}
	}

	return aiResponse, nil
}

// Utility functions for AI processing

// ValidateTemplate checks if a template has all required variables
func (ai *AICore) ValidateTemplate(templateID string, variables map[string]interface{}) error {
	template, exists := ai.PromptTemplates[templateID]
	if !exists {
		return fmt.Errorf("template not found: %s", templateID)
	}

	// Check if all required variables are provided
	for _, variable := range template.Variables {
		if _, exists := variables[variable]; !exists {
			return fmt.Errorf("missing required variable: %s", variable)
		}
	}

	return nil
}

// GetPersonalityByLevel returns available personalities for a user level
func (ai *AICore) GetPersonalityByLevel(userLevel int) []BuddyPersonality {
	var available []BuddyPersonality
	
	for personality, config := range ai.PersonalityMap {
		if userLevel >= config.UnlockLevel {
			available = append(available, personality)
		}
	}
	
	return available
}

// CalculateResponseQuality estimates the quality of an AI response
func (ai *AICore) CalculateResponseQuality(response string, expectedLength int) float64 {
	// Simple quality metrics
	length := len(response)
	
	// Length score (0-1)
	lengthScore := 1.0
	if expectedLength > 0 {
		ratio := float64(length) / float64(expectedLength)
		if ratio < 0.5 || ratio > 2.0 {
			lengthScore = 0.5
		}
	}
	
	// Content quality indicators
	qualityScore := 0.5 // Base score
	
	// Check for helpful indicators
	if strings.Contains(strings.ToLower(response), "step") {
		qualityScore += 0.1
	}
	if strings.Contains(strings.ToLower(response), "help") {
		qualityScore += 0.1
	}
	if strings.Contains(strings.ToLower(response), "learn") {
		qualityScore += 0.1
	}
	if len(strings.Split(response, ".")) > 3 { // Multiple sentences
		qualityScore += 0.1
	}
	
	// Ensure score is between 0 and 1
	if qualityScore > 1.0 {
		qualityScore = 1.0
	}
	
	return (lengthScore + qualityScore) / 2.0
}

