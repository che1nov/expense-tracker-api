package handlers

import (
	"encoding/json"
	database "expense-tracker-api/internal/db"
	"expense-tracker-api/internal/models"

	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		slog.Error("Invalid request payload", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		slog.Error("Failed to create user", slog.String("username", user.Username))
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	slog.Info("User signed up successfully", slog.String("username", user.Username))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Signup successful"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Error("Invalid request payload", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Where("username = ?", input.Username).First(&user)
	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(72 * time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
