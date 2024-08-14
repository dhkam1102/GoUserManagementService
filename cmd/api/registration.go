package api

import (
	"encoding/json"
	"log"
	"net/http"

	"user-management-service/internal/database"
	"user-management-service/internal/models"
)

func RegisterHandler(writer http.ResponseWriter, request *http.Request) {
	// Logging the start of the request -------
	log.Printf("Received a registration request from %s", request.RemoteAddr)

	// NOTE: creating a RegisterRequest struct to hold the data
	var req models.RegisterRequest

	// Logging the raw request body for debugging --------
	log.Printf("Request body: %v", request.Body)

	// NOTE: creating a json decoder that will read request's body
	decoder := json.NewDecoder(request.Body)
	// NOTE: decode json into req
	error := decoder.Decode(&req)

	if error != nil {
		http.Error(writer, "invalid request type1", http.StatusBadRequest)
		return
	}

	log.Printf("Decoded registration request: Email: %s, Password: %s, ConfirmPassword: %s", req.Email, req.Password, req.ConfirmPass)

	// IMPROVE: might want to use hashing for security

	if req.Password != req.ConfirmPass {
		http.Error(writer, "non matching password and confirm password", http.StatusBadRequest)
		return
	}

	if req.Role == "" {
		req.Role = "customer"
	}

	user := models.User{
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	log.Printf("Connecting to database to insert user: %s", user.Email)

	db := database.NewDatabaseConnection()
	// NOTE: ensures db connection is closed when function end
	defer db.Close()

	// Logging before executing the insert statement
	log.Printf("Attempting to insert user: %s into the database", user.Email)

	// NOTE: db.Exec return value: sql.Result(RowsAffected(), LastInsertId()), error
	_, error = db.Exec("INSERT INTO users (email, password, role) VALUES (?, ?, ?)", user.Email, user.Password, user.Role)
	if error != nil {
		http.Error(writer, "Error on saving new user", http.StatusInternalServerError)
		return
	}

	// Logging successful registration
	log.Printf("User %s registered successfully", user.Email)

	writer.WriteHeader(http.StatusCreated)
	// NOTE: []byte converting strings into slices of bytes
	// NOTE: by slicing the string into bytes, it can be sent as HTTP response body
	writer.Write([]byte("User registered successfully"))
}

//fklajvaljval;jvd testing
