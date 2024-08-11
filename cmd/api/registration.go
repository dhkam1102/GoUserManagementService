package api

import (
	"encoding/json"
	"net/http"
	"user-management-service/internal/database"
	"user-management-service/internal/models"
)

type RegisterRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	ConfirmPass string `json:"confirm_password"`
}

func RegisterHandler(writer http.ResponseWriter, request *http.Request) {
	// NOTE: creating a RegisterRequest struct to hold the data
	var req RegisterRequest
	// NOTE: creating a json decoder that will read request's body
	decoder := json.NewDecoder(request.Body)
	// NOTE: decode json into req
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

	user := models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	db := database.NewDatabaseConnection()
	// NOTE: ensures db connection is closed when function end
	defer db.Close()
	// NOTE: db.Exec return value: sql.Result(RowsAffected(), LastInsertId()), error
	_, error = db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", user.Email, user.Password)
	if error != nil {
		http.Error(writer, "Error on saving new user", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	// NOTE: []byte converting strings into slices of bytes
	// NOTE: by slicing the string into bytes, it can be sent as HTTP response body
	writer.Write([]byte("User registered successfully"))
}
