package backend

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	public := router.Group("/api")
	{
		public.POST("/login", handleLogin)
		public.POST("/register", handleRegister)
	}

	protected := router.Group("/api")
	protected.Use(authMiddleware())
	{
		protected.GET("/questions", handleGetQuestions)

		// User profile
		protected.GET("/profile", handleGetProfile)
		protected.PUT("/profile", handleUpdateProfile)

		// Questionnaire submissions
		protected.POST("/submit", handleSubmitQuestionnaire)

		// Admin routes
		admin := protected.Group("/admin")
		admin.Use(adminMiddleware())
		{
			admin.GET("/submissions", handleGetUserSubmissions)
			admin.GET("/submissions/:id", handleGetSubmission)
			admin.POST("/questions", handleCreateQuestion)
			admin.PUT("/questions/:id", handleUpdateQuestion)
			admin.DELETE("/questions/:id", handleDeleteQuestion)
			admin.GET("/users", handleGetAllUsers)
			admin.GET("/submissions/all", handleGetAllSubmissions)
		}
	}
}

func handleLogin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericResponse{Success: false, Data: err.Error()})
		return
	}

	user, err := db.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, GenericResponse{Success: false, Data: "Invalid credentials"})
		return
	}

	if !checkPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, GenericResponse{Success: false, Data: "Invalid credentials"})
		return
	}

	token, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Success: false, Data: "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, GenericResponse{
		Success: true,
		Data: map[string]string{
			"token": token,
		},
	})
}

func handleRegister(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericResponse{Success: false, Data: err.Error()})
		return
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Success: false, Data: "Error processing password"})
		return
	}

	user := &User{
		Email:    req.Email,
		Password: hashedPassword,
		Name:     req.Name,
		IsAdmin:  false,
	}

	if err := db.CreateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, GenericResponse{Success: false, Data: err.Error()})
		return
	}

	token, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Success: false, Data: "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, GenericResponse{
		Success: true,
		Data: map[string]string{
			"token": token,
		},
	})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, GenericResponse{Success: false, Data: "Authorization required"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := validateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, GenericResponse{Success: false, Data: "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("isAdmin", claims.IsAdmin)
		c.Next()
	}
}

func adminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, exists := c.Get("isAdmin")
		if !exists || !isAdmin.(bool) {
			c.JSON(http.StatusForbidden, GenericResponse{Success: false, Data: "Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func handleGetProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	user, err := db.GetUser(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, GenericResponse{Success: false, Data: "User not found"})
		return
	}

	c.JSON(http.StatusOK, GenericResponse{
		Success: true,
		Data: ProfileResponse{
			ID:      user.ID,
			Name:    user.Name,
			Picture: user.Picture,
			IsAdmin: user.IsAdmin,
		},
	})
}

func handleUpdateProfile(c *gin.Context) {
	userID, _ := c.Get("userID")
	isAdmin, _ := c.Get("isAdmin")

	// Check authorization
	if userID == nil || userID.(string) == "" || (userID.(string) != c.Param("id") && isAdmin != true) {
		c.JSON(http.StatusUnauthorized, GenericResponse{Success: false, Data: "Unauthorized"})
		return
	}

	var updateReq struct {
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, GenericResponse{Success: false, Data: err.Error()})
		return
	}

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(usersBucket)

		userData := b.Get([]byte(userID.(string)))
		if userData == nil {
			return errors.New("user not found")
		}

		var user User
		if err := json.Unmarshal(userData, &user); err != nil {
			return err
		}

		user.Name = updateReq.Name
		user.Picture = updateReq.Picture

		buf, err := json.Marshal(user)
		if err != nil {
			return err
		}

		return b.Put([]byte(userID.(string)), buf)
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Success: false, Data: "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, GenericResponse{Success: true, Data: "Profile updated successfully"})
}

func handleGetQuestions(c *gin.Context) {
	questions, err := db.GetQuestions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Success: false, Data: "Failed to fetch questions"})
		return
	}

	c.JSON(http.StatusOK, GenericResponse{
		Success: true,
		Data: QuestionsResponse{
			Questions: questions,
		},
	})
}

func handleCreateQuestion(c *gin.Context) {
	var question Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, GenericResponse{Success: false, Data: err.Error()})
		return
	}

	validType := false
	for _, t := range questionTypes {
		if question.Type == t {
			validType = true
			break
		}
	}
	if !validType {
		c.JSON(http.StatusBadRequest, GenericResponse{Success: false, Data: "Invalid question type"})
		return
	}

	if question.Type == "scale" && (question.Min == 0 || question.Max == 0) {
		c.JSON(http.StatusBadRequest, GenericResponse{Success: false, Data: "Scale questions require min and max values"})
		return
	}

	if err := db.CreateQuestion(&question); err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Success: false, Data: "Failed to create question"})
		return
	}

	c.JSON(http.StatusOK, GenericResponse{Success: true, Data: question})
}

func handleUpdateQuestion(c *gin.Context) {
	questionID := c.Param("id")

	var updateReq Question
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, GenericResponse{Success: false, Data: err.Error()})
		return
	}

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(questionsBucket)

		existing := b.Get([]byte(questionID))
		if existing == nil {
			return errors.New("question not found")
		}

		updateReq.ID = questionID

		validType := false
		for _, t := range questionTypes {
			if updateReq.Type == t {
				validType = true
				break
			}
		}
		if !validType {
			return errors.New("invalid question type")
		}

		buf, err := json.Marshal(updateReq)
		if err != nil {
			return err
		}
		return b.Put([]byte(questionID), buf)
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Success: false, Data: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GenericResponse{Success: true, Data: updateReq})
}

func handleDeleteQuestion(c *gin.Context) {
	questionID := c.Param("id")

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(questionsBucket)

		if existing := b.Get([]byte(questionID)); existing == nil {
			return errors.New("question not found")
		}

		return b.Delete([]byte(questionID))
	})

	if err != nil {
		if err.Error() == "question not found" {
			c.JSON(http.StatusNotFound, GenericResponse{Success: false, Data: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, GenericResponse{Success: false, Data: "Failed to delete question"})
		}
		return
	}

	c.JSON(http.StatusOK, GenericResponse{Success: true, Data: "Question deleted successfully"})
}

func handleSubmitQuestionnaire(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req AnswersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericResponse{Success: false, Data: err.Error()})
		return
	}

	questions, err := db.GetQuestions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Success: false, Data: "Failed to validate questions"})
		return
	}

	questionMap := make(map[string]Question)
	for _, q := range questions {
		questionMap[q.ID] = q
	}

	for _, answer := range req.Answers {
		if _, exists := questionMap[answer.ID]; !exists {
			c.JSON(http.StatusBadRequest, GenericResponse{Success: false, Data: "Invalid question ID: " + answer.ID})
			return
		}
	}

	submission := &Submission{
		UserID:  userID.(string),
		Answers: req.Answers,
	}

	if err := db.CreateSubmission(submission); err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Success: false, Data: "Failed to save submission"})
		return
	}

	c.JSON(http.StatusOK, GenericResponse{
		Success: true,
		Data: SubmissionResponse{
			ID:        submission.ID,
			CreatedAt: submission.CreatedAt,
		},
	})
}

func handleGetSubmission(c *gin.Context) {
	userID, _ := c.Get("userID")
	isAdmin, _ := c.Get("isAdmin")
	submissionID := c.Param("id")

	var submission Submission
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(submissionsBucket)
		data := b.Get([]byte(submissionID))
		if data == nil {
			return errors.New("submission not found")
		}
		return json.Unmarshal(data, &submission)
	})

	if err != nil {
		c.JSON(http.StatusNotFound, GenericResponse{Success: false, Data: "Submission not found"})
		return
	}

	if !isAdmin.(bool) && submission.UserID != userID.(string) {
		c.JSON(http.StatusForbidden, GenericResponse{Success: false, Data: "Access denied"})
		return
	}

	c.JSON(http.StatusOK, GenericResponse{Success: true, Data: submission})
}

func handleGetUserSubmissions(c *gin.Context) {
	userID, _ := c.Get("userID")

	submissions, err := db.GetUserSubmissions(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Success: false, Data: "Failed to fetch submissions"})
		return
	}

	c.JSON(http.StatusOK, GenericResponse{
		Success: true,
		Data: UserSubmissionsResponse{
			Submissions: submissions,
		},
	})
}

func handleGetAllUsers(c *gin.Context) {
	users, err := db.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Success: false, Data: "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, GenericResponse{
		Success: true,
		Data: UsersResponse{
			Users: users,
		},
	})
}

func handleGetAllSubmissions(c *gin.Context) {
	submissions, err := db.GetAllSubmissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Success: false, Data: "Failed to fetch submissions"})
		return
	}

	if submissions == nil {
		submissions = []Submission{}
	}

	c.JSON(http.StatusOK, GenericResponse{
		Success: true,
		Data: UserSubmissionsResponse{
			Submissions: submissions,
		},
	})
}
