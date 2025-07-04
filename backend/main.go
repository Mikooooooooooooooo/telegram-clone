package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func generateToken(phone string) (string, error) {
	claims := jwt.MapClaims{
		"phone": phone,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}

func authMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	token, err := validateToken(authHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Set("phone", claims["phone"])
	c.Next()
}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Phone    string `gorm:"unique"`
	Password string
}

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=12345678 dbname=telegram_clone port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Println("✅ Connected to database")

	db.AutoMigrate(&User{})
}

func main() {

	initDB()

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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r.POST("/check-user", func(c *gin.Context) {
		var body struct {
			Phone string `json:"phone"`
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Output"})
			return
		}

		var user User
		result := db.Where("phone = ?", body.Phone).First(&user)

		c.JSON(http.StatusOK, gin.H{"exists": result.Error == nil})
	})

	r.POST("/register", func(c *gin.Context) {
		var body struct {
			Phone    string `json:"phone"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
			return
		}
		var existing User
		if err := db.Where("phone = ? ", body.Phone).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
			return
		}

		newUser := User{
			Phone:    body.Phone,
			Password: string(hashedPassword),
		}

		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

	})

	r.POST("/login", func(c *gin.Context) {
		var body struct {
			Phone    string `json:"phone"`
			Password string `json:"password"`
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
			return
		}

		var user User
		if err := db.Where("phone = ?", body.Phone).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
			return
		}

		token, err := generateToken(user.Phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"token":   token,
		})
	})

	r.GET("/me", authMiddleware, func(c *gin.Context) {
		phone := c.MustGet("phone").(string)
		c.JSON(http.StatusOK, gin.H{"message": "Welcome back!", "phone": phone})
	})

	r.Run(":8080")
}
