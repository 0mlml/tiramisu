package backend

import "time"

// GenericResponse represents a generic JSON response
type GenericResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ProfileResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	IsAdmin bool   `json:"is_admin"`
	Created string `json:"created"`
}

type Answer struct {
	ID       string `json:"id"`
	Question string `json:"question"`
}

type AnswersRequest struct {
	Answers []Answer `json:"answers"`
}

var questionTypes = []string{"scale"}

type Question struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	Type     string `json:"type"`
	Min      int    `json:"min,omitempty"`
	Max      int    `json:"max,omitempty"`
}

type QuestionsResponse struct {
	Questions []Question `json:"questions"`
}

// User represents a user in the system
type User struct {
	ID       string    `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Picture  string    `json:"picture"`
	IsAdmin  bool      `json:"is_admin"`
	Created  time.Time `json:"created"`
}

// LoginRequest represents the login form data
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest represents the registration form data
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

// Submission represents a completed questionnaire
type Submission struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Answers   []Answer  `json:"answers"`
	CreatedAt time.Time `json:"created_at"`
}

// SubmissionResponse represents the response for a questionnaire submission
type SubmissionResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

// UserSubmissionsResponse represents a list of user submissions
type UserSubmissionsResponse struct {
	Submissions []Submission `json:"submissions"`
}

// UsersResponse represents a list of users
type UsersResponse struct {
	Users []User `json:"users"`
}
