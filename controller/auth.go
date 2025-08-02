package controllers

import (
	"encoding/json"
	"main/config"
	"main/models"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Basic validations
	if user.FirstName == "" {
		http.Error(w, "First name is required", http.StatusBadRequest)
		return
	}
	if user.LastName == "" {
		http.Error(w, "Last name is required", http.StatusBadRequest)
		return
	}
	if user.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	if user.Email != "" {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}
	if len(user.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters long", http.StatusBadRequest)
		return
	}

	// Check if user already exists
	var exists int
	err = config.DB.QueryRow("SELECT 1 FROM Users WHERE Email = @p1", user.Email).Scan(&exists)
	if err == nil {
		response := map[string]interface{}{
			"message": "User already registered",
			"status":  http.StatusConflict,
			"error":   true,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	} else if err.Error() != "sql: no rows in result set" {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Insert into database
	insertQuery := `INSERT INTO Users (FirstName, LastName, Email, PasswordHash) VALUES (@p1, @p2, @p3, @p4)`
	_, err = config.DB.Exec(insertQuery, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	response := map[string]interface{}{
		"message": "User registered successfully",
		"user": map[string]string{
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"email":     user.Email,
		},
		"status": http.StatusCreated,
		"error":  false,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
