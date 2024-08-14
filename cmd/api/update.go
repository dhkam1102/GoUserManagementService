package api

import (
	"encoding/json"
	"log"
	"net/http"
	"user-management-service/internal/database"
	"user-management-service/internal/models"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.NewEmail == "" && req.NewPassword == "" {
		http.Error(w, "No fields to update", http.StatusBadRequest)
		return
	}

	db := database.NewDatabaseConnection()
	defer db.Close()

	get_id_query := "SELECT id FROM users WHERE email = ? AND password = ?"
	var userID int
	err = db.QueryRow(get_id_query, req.OldEmail, req.OldPassword).Scan(&userID)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	updateQuery := "UPDATE users SET email = ?, password = ? WHERE id = ?"
	_, err = db.Exec(updateQuery, req.NewEmail, req.NewPassword, userID)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	log.Printf("New email %s updated successfully", req.NewEmail)

	// should be the token part
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "User updated successfully"}`))
}
