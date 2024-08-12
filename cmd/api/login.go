package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"user-management-service/internal/database"
	"user-management-service/internal/models"
)

// NOTE: LoginRequest has no confirm password
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	log.Printf("LoginHandler: Received login request from IP: %s", r.RemoteAddr)

	// NOTE: creating a json decoder that will read request's body
	decoder := json.NewDecoder(r.Body)
	// NOTE: decode json into req
	error := decoder.Decode(&req)

	if error != nil {
		log.Printf("Lofing handler Error in decoding: %v", error)
		http.Error(w, "invalid request type1", http.StatusBadRequest)
		return
	}

	// Logging before attempting database connection
	log.Printf("Connecting to database to insert user: %s", req.Email)

	db := database.NewDatabaseConnection()
	defer db.Close()

	log.Printf("LoginHandler: Database connection established")

	var user models.User
	err := db.QueryRow("SELECT id, email, password FROM users WHERE email = ?", req.Email).Scan(&user.ID, &user.Email, &user.Password)

	// NOTE: check if the email does not exist
	if err == sql.ErrNoRows {
		// Email not found in the database
		log.Printf("LoginHandler: Email not found - Email: %s", req.Email)
		http.Error(w, "Invalid email", http.StatusUnauthorized)
		return
	}

	// NOTE: Handle other potential errors (database connection, query error)
	if err != nil {
		log.Printf("LoginHandler: Error querying database: %v", err)
		http.Error(w, "Something went wrong. Please try again later.", http.StatusInternalServerError)
		return
	}

	log.Printf("LoginHandler: User found - ID: %d, Email: %s", user.ID, user.Email)
	// IMPORVE: use hashed password

	// NOTE: need more understanding of token and the lines below
	// NOTE: jwt token has 3parts: hearder, payload, signature
	// NOTE: using token we can have a stateless authentication
	token := "jwt-token"

	// NOTE: sending the token back to user
	header := w.Header()
	// NOTE: setting content-type
	header.Set("Content-Type", "application/json")
	// Create a map to hold the JSON data
	responseData := make(map[string]string)

	// Set the "token" key to the value of the token variable
	responseData["token"] = token
	// Create a new JSON encoder that writes to the ResponseWriter
	encoder := json.NewEncoder(w)

	// Encode the map into JSON format and write it to the response
	err = encoder.Encode(responseData)

	// Check for errors in the encoding process
	if err != nil {
		// Handle the error, possibly by sending an internal server error status
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
