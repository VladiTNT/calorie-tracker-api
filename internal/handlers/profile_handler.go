package handlers

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/VladiTNT/calorie-tracker-api/internal/database/models"
	"github.com/VladiTNT/calorie-tracker-api/internal/database/queries"
	"github.com/VladiTNT/calorie-tracker-api/pkg/auth"
)

type ProfileHandler struct {
	Logger   *slog.Logger
	Database *sql.DB
}

func NewProfileHandler(l *slog.Logger, db *sql.DB) *ProfileHandler {
	return &ProfileHandler{
		l, db,
	}
}

func (ph *ProfileHandler) Err(w http.ResponseWriter, err error, statusCode int) {
	// Log error
	ph.Logger.Error(err.Error(), "status", statusCode)

	// Json error message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]any{
		"status": statusCode,
		"error":  err.Error(),
	})
}

func (ph *ProfileHandler) GetProfileSelf(w http.ResponseWriter, r *http.Request) {
	name, ok := r.Context().Value(auth.AuthContextKey).(string)
	if !ok {
		ph.Err(w, ErrUserNotAuthenticated, http.StatusUnauthorized)
		return
	}

	p, err := queries.GetProfileByName(ph.Database, r.Context(), name)
	if err != nil {
		ph.Err(w, err, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		ph.Err(w, err, http.StatusInternalServerError)
	}
}

func (ph *ProfileHandler) GetProfileByName(w http.ResponseWriter, r *http.Request) {
	p, err := queries.GetProfileByName(ph.Database, r.Context(), r.PathValue("name"))
	if err != nil {
		ph.Err(w, err, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		ph.Err(w, err, http.StatusInternalServerError)
	}
}

func (ph *ProfileHandler) UpdateProfileCalories(w http.ResponseWriter, r *http.Request) {
	name, ok := r.Context().Value(auth.AuthContextKey).(string)
	if !ok {
		ph.Err(w, ErrUserNotAuthenticated, http.StatusUnauthorized)
		return
	}

	var p models.UserProfile
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		ph.Err(w, err, http.StatusBadRequest)
		return
	}

	err := queries.UpdateProfileCalories(ph.Database, r.Context(), name, p.TargetCalories)
	if err != nil {
		ph.Err(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"status":  http.StatusOK,
		"message": "profile calories updated",
	})
}
