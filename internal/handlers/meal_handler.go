package handlers

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/VladiTNT/calorie-tracker-api/internal/database/models"
	"github.com/VladiTNT/calorie-tracker-api/internal/database/queries"
	"github.com/VladiTNT/calorie-tracker-api/pkg/auth"
)

type MealHandler struct {
	Logger   *slog.Logger
	Database *sql.DB
}

func NewMealHandler(l *slog.Logger, db *sql.DB) *MealHandler {
	return &MealHandler{l, db}
}

func (mh *MealHandler) Err(w http.ResponseWriter, err error, statusCode int) {
	// Log error
	mh.Logger.Error(err.Error(), "status", statusCode)

	// Json error message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]any{
		"status": statusCode,
		"error":  err.Error(),
	})
}

func (mh *MealHandler) PostMeal(w http.ResponseWriter, r *http.Request) {
	name, ok := r.Context().Value(auth.AuthContextKey).(string)
	if !ok {
		mh.Err(w, ErrUserNotAuthenticated, http.StatusUnauthorized)
		return
	}

	var m models.Meal
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		mh.Err(w, err, http.StatusBadRequest)
		return
	}

	m.User = name

	if err := queries.InsertMeal(mh.Database, r.Context(), &m); err != nil {
		mh.Err(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"status":  http.StatusCreated,
		"message": "meal added to database",
	})
}

func (mh *MealHandler) GetMealsToday(w http.ResponseWriter, r *http.Request) {
	name, ok := r.Context().Value(auth.AuthContextKey).(string)
	if !ok {
		mh.Err(w, ErrUserNotAuthenticated, http.StatusUnauthorized)
		return
	}

	meals, err := queries.GetMealsOnDayByUsername(mh.Database, r.Context(), name, time.Now())
	if err != nil {
		mh.Err(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(meals); err != nil {
		mh.Err(w, err, http.StatusInternalServerError)
		return
	}
}

func (mh *MealHandler) GetMealsTodayUser(w http.ResponseWriter, r *http.Request) {
	meals, err := queries.GetMealsOnDayByUsername(mh.Database, r.Context(), r.PathValue("user"), time.Now())
	if err != nil {
		mh.Err(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(meals); err != nil {
		mh.Err(w, err, http.StatusInternalServerError)
		return
	}
}

func (mh *MealHandler) GetMealsDateUser(w http.ResponseWriter, r *http.Request) {
	t, err := time.Parse(time.DateOnly, r.PathValue("day"))
	if err != nil {
		mh.Err(w, err, http.StatusBadRequest)
		return
	}

	meals, err := queries.GetMealsOnDayByUsername(mh.Database, r.Context(), r.PathValue("user"), t)
	if err != nil {
		mh.Err(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(meals); err != nil {
		mh.Err(w, err, http.StatusInternalServerError)
		return
	}
}
