package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"user-management-service/internal/database"
	"user-management-service/internal/models"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key") // Use a secure key

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&req)

	if error != nil {
		log.Printf("Lofing handler Error in decoding: %v", error)
		http.Error(w, "invalid request type1", http.StatusBadRequest)
		return
	}

	db := database.NewDatabaseConnection()
	defer db.Close()

	var user models.User
	err := db.QueryRow("SELECT id, email, password FROM users WHERE email = ?", req.Email).Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		log.Printf("LoginHandler: Email not found - Email: %s", req.Email)
		http.Error(w, "Invalid email", http.StatusUnauthorized)
		return
	}

	if err != nil {
		log.Printf("LoginHandler: Error querying database: %v", err)
		http.Error(w, "Something went wrong. Please try again later.", http.StatusInternalServerError)
		return
	}

	// IMPORVE: use hashed password
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: req.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("LoginHandler: Error generating JWT token: %v", err)
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	// Return the JWT token
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token": "` + tokenString + `"}`))

	log.Printf("LoginHandler: Successfully sent response with JWT token")
}
