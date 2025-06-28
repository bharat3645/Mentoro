package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// User represents a user in the system
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Level    int    `json:"level"`
	XP       int    `json:"xp"`
	Streak   int    `json:"streak"`
	Mood     string `json:"mood"`
}

// Quest represents a learning quest
type Quest struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Progress    int    `json:"progress"`
	Total       int    `json:"total"`
	XP          int    `json:"xp"`
	Type        string `json:"type"`
	UserID      int    `json:"user_id"`
}

// Badge represents an achievement badge
type Badge struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Earned      bool   `json:"earned"`
	UserID      int    `json:"user_id"`
}

// BuddyMessage represents AI buddy conversation
type BuddyMessage struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Message   string    `json:"message"`
	Response  string    `json:"response"`
	Mood      string    `json:"mood"`
	Timestamp time.Time `json:"timestamp"`
}

// Mock data for demonstration
var mockUsers = []User{
	{ID: 1, Username: "alex", Email: "alex@example.com", Level: 5, XP: 750, Streak: 7, Mood: "focused"},
}

var mockQuests = []Quest{
	{ID: 1, Title: "Fix 3 bugs", Description: "Debug and fix 3 code issues", Progress: 2, Total: 3, XP: 150, Type: "code", UserID: 1},
	{ID: 2, Title: "10 min focus session", Description: "Complete a focused work session", Progress: 7, Total: 10, XP: 50, Type: "focus", UserID: 1},
	{ID: 3, Title: "Learn new concept", Description: "Study and understand a new programming concept", Progress: 0, Total: 1, XP: 200, Type: "learn", UserID: 1},
}

var mockBadges = []Badge{
	{ID: 1, Name: "Code Warrior", Description: "Fixed 10 bugs", Icon: "⚔️", Earned: true, UserID: 1},
	{ID: 2, Name: "Focus Master", Description: "Completed 50 focus sessions", Icon: "🎯", Earned: true, UserID: 1},
	{ID: 3, Name: "Streak Champion", Description: "Maintained 30-day streak", Icon: "🔥", Earned: false, UserID: 1},
	{ID: 4, Name: "Bug Hunter", Description: "Found and reported 5 bugs", Icon: "🐛", Earned: true, UserID: 1},
}

var mockMessages = []BuddyMessage{
	{ID: 1, UserID: 1, Message: "I'm feeling stuck on this problem", Response: "I understand! Let's break it down into smaller steps. What specific part is challenging you?", Mood: "mentor", Timestamp: time.Now()},
}

// API Handlers

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	for _, user := range mockUsers {
		if user.ID == userID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func getUserQuestsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var userQuests []Quest
	for _, quest := range mockQuests {
		if quest.UserID == userID {
			userQuests = append(userQuests, quest)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userQuests)
}

func getUserBadgesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var userBadges []Badge
	for _, badge := range mockBadges {
		if badge.UserID == userID {
			userBadges = append(userBadges, badge)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userBadges)
}

func updateQuestProgressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid quest ID", http.StatusBadRequest)
		return
	}

	var updateData struct {
		Progress int `json:"progress"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, quest := range mockQuests {
		if quest.ID == questID {
			mockQuests[i].Progress = updateData.Progress
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(mockQuests[i])
			return
		}
	}

	http.Error(w, "Quest not found", http.StatusNotFound)
}

func buddyChatHandler(w http.ResponseWriter, r *http.Request) {
	var chatRequest struct {
		UserID  int    `json:"user_id"`
		Message string `json:"message"`
		Mood    string `json:"mood"`
	}

	if err := json.NewDecoder(r.Body).Decode(&chatRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Simple mock AI response based on mood
	responses := map[string][]string{
		"mentor": {
			"Let me help you understand this concept better. What specific part would you like me to explain?",
			"Great question! This is a common challenge. Let's approach it step by step.",
			"I can see you're working hard. Remember, learning is a process - take your time.",
		},
		"cheerleader": {
			"You're doing amazing! Keep up the great work! 🎉",
			"I believe in you! You've got this! 💪",
			"Every step forward is progress. You're crushing it! ⭐",
		},
		"chill": {
			"No worries, we'll figure this out together. Take a deep breath.",
			"Hey, no pressure. Let's just explore this at your own pace.",
			"All good! Sometimes the best solutions come when we're relaxed.",
		},
		"focused": {
			"Let's dive deep into this problem. What data do we have?",
			"Time to focus. What's the core issue we need to solve?",
			"Good, you're in the zone. Let's tackle this systematically.",
		},
	}

	mood := chatRequest.Mood
	if mood == "" {
		mood = "mentor"
	}

	responseList, exists := responses[mood]
	if !exists {
		responseList = responses["mentor"]
	}

	// Simple response selection (in real app, this would use OpenAI)
	response := responseList[len(chatRequest.Message)%len(responseList)]

	buddyResponse := BuddyMessage{
		ID:        len(mockMessages) + 1,
		UserID:    chatRequest.UserID,
		Message:   chatRequest.Message,
		Response:  response,
		Mood:      mood,
		Timestamp: time.Now(),
	}

	mockMessages = append(mockMessages, buddyResponse)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buddyResponse)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"service":   "learning-buddy-backend",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	// API Routes
	api := r.PathPrefix("/api/v1").Subrouter()
	
	// Health check
	api.HandleFunc("/health", healthCheckHandler).Methods("GET")
	
	// User routes
	api.HandleFunc("/users/{id}", getUserHandler).Methods("GET")
	api.HandleFunc("/users/{id}/quests", getUserQuestsHandler).Methods("GET")
	api.HandleFunc("/users/{id}/badges", getUserBadgesHandler).Methods("GET")
	
	// Quest routes
	api.HandleFunc("/quests/{id}/progress", updateQuestProgressHandler).Methods("PUT")
	
	// AI Buddy routes
	api.HandleFunc("/buddy/chat", buddyChatHandler).Methods("POST")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(r)

	port := "8080"
	fmt.Printf("🚀 Learning Buddy Backend starting on port %s\n", port)
	fmt.Printf("📊 Health check: http://localhost:%s/api/v1/health\n", port)
	fmt.Printf("👤 User API: http://localhost:%s/api/v1/users/1\n", port)
	fmt.Printf("🎯 Quests API: http://localhost:%s/api/v1/users/1/quests\n", port)
	fmt.Printf("🏆 Badges API: http://localhost:%s/api/v1/users/1/badges\n", port)
	fmt.Printf("🤖 Buddy Chat: http://localhost:%s/api/v1/buddy/chat\n", port)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, handler))
}

