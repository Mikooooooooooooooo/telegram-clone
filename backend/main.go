package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type User struct {
	Phone    string
	Password string
}

var users = map[string]User{}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/check-user", func(c *gin.Context) {
		var body struct {
			Phone string `json:"phone"`
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Output"})
			return
		}

		_, exists := users[body.Phone]
		c.JSON(http.StatusOK, gin.H{"exists": exists})
	})

	r.POST("/register", func(c *gin.Context) {
		var body struct {
			Phone    string `json: "phone"`
			Password string `json: "password"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
			return
		}
		if _, exists := users[body.Phone]; exists {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
			return
		}

		users[body.Phone] = User{
			Phone:    body.Phone,
			Password: string(hashedPassword),
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

	})

	r.POST("/login", func(c *gin.Context) {
		var body struct {
			Phone    string `json: "phone"`
			Password string `json: "password"`
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
			return
		}
		user, exists := users[body.Phone]

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})

	r.Run(":8080")
}
