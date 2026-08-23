package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/VladiTNT/calorie-tracker-api/internal/database/models"
	"github.com/VladiTNT/calorie-tracker-api/internal/database/queries"
	"github.com/VladiTNT/calorie-tracker-api/pkg/auth"
)

var ErrUserNotAuthenticated = errors.New("User does not bear an authentication token.")

type AuthHandler struct {
	Logger      *slog.Logger
	AuthService *auth.Service
	Database    *sql.DB
}

func NewAuthHandler(l *slog.Logger, as *auth.Service, db *sql.DB) *AuthHandler {
	return &AuthHandler{l, as, db}
}

func (ah *AuthHandler) Err(w http.ResponseWriter, err error, statusCode int) {
	// Log error
	ah.Logger.Error(err.Error(), "status", statusCode)

	// Write the status code and send an error json encoded message to the client
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status": statusCode,
		"error":  err.Error(),
	})
}

func (ah *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		ah.Err(w, err, http.StatusBadRequest)
		return
	}

	// Database transaction
	tx, err := ah.Database.BeginTx(r.Context(), nil)
	if err != nil {
		ah.Err(w, err, http.StatusInternalServerError)
		return
	}

	// Insert user credentials
	if err := queries.InsertUser(tx, r.Context(), &u); err != nil {
		ah.Err(w, err, http.StatusInternalServerError)
		return
	}

	// Create a profile for the new user
	var p models.UserProfile = models.UserProfile{Name: u.Name, TargetCalories: 0}
	if err := queries.InsertProfile(tx, r.Context(), &p); err != nil {
		ah.Err(w, err, http.StatusInternalServerError)
		return
	}

	// Commit database transaction
	if err := tx.Commit(); err != nil {
		ah.Err(w, err, http.StatusInternalServerError)
		return
	}

	// Pass the client an auth token
	c := ah.AuthService.Add(u.Name)
	http.SetCookie(w, &c)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  http.StatusCreated,
		"message": "user created",
	})
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		ah.Err(w, err, http.StatusBadRequest)
		return
	}

	isValid := queries.ValidateUser(ah.Database, r.Context(), &u)

	w.Header().Set("Content-Type", "application/json")

	if isValid {
		c := ah.AuthService.Add(u.Name)
		http.SetCookie(w, &c)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  http.StatusOK,
			"message": "user authenticated",
		})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  http.StatusUnauthorized,
			"message": "user password or name is invalid",
		})
	}
}
