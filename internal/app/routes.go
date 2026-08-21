package app

import "github.com/VladiTNT/calorie-tracker-api/internal/handlers"

func (ctapi *CalorieTrackerAPI) RegisterRoutes() {
	// Auth
	authHandler := handlers.NewAuthHandler(ctapi.Logger, ctapi.AuthService, ctapi.Database)

	ctapi.Router.HandleFunc("POST /api/auth/signup", authHandler.Signup)
	ctapi.Router.HandleFunc("POST /api/auth/login", authHandler.Login)
}
