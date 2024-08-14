package api

import (
	"encoding/json"
	// "log"
	"net/http"

	"user-management-service/internal/database"
	"user-management-service/internal/models"
)

func RegisterHandler(writer http.ResponseWriter, request *http.Request) {

	var req models.RegisterRequest
	decoder := json.NewDecoder(request.Body)
	error := decoder.Decode(&req)
	if error != nil {
		http.Error(writer, "invalid request type1", http.StatusBadRequest)
		return
	}

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

	db := database.NewDatabaseConnection()
	defer db.Close()

	_, error = db.Exec("INSERT INTO users (email, password, role) VALUES (?, ?)", user.Email, user.Password, user.Role)
	if error != nil {
		http.Error(writer, "Error on saving new user", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("User registered successfully"))
}
