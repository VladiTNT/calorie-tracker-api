package app

import "github.com/VladiTNT/calorie-tracker-api/internal/handlers"

func (ctapi *CalorieTrackerAPI) RegisterRoutes() {
	// Auth
	authHandler := handlers.NewAuthHandler(ctapi.Logger, ctapi.AuthService, ctapi.Database)

	ctapi.Router.HandleFunc("POST /api/auth/signup", authHandler.Signup)
	ctapi.Router.HandleFunc("POST /api/auth/login", authHandler.Login)

	// Profiles
	profileHandler := handlers.NewProfileHandler(ctapi.Logger, ctapi.Database)

	ctapi.Router.HandleFunc("GET /api/profile", profileHandler.GetProfileSelf)
	ctapi.Router.HandleFunc("GET /api/profile/{name}", profileHandler.GetProfileByName)
	ctapi.Router.HandleFunc("PUT /api/profile/target_calories", profileHandler.UpdateProfileCalories)

	// Meals
	mealHandler := handlers.NewMealHandler(ctapi.Logger, ctapi.Database)

	ctapi.Router.HandleFunc("POST /api/meal", mealHandler.PostMeal)
	ctapi.Router.HandleFunc("GET /api/meal", mealHandler.GetMealsToday)
	ctapi.Router.HandleFunc("GET /api/meal/{user}", mealHandler.GetMealsTodayUser)
	ctapi.Router.HandleFunc("GET /api/meal/{user}/{day}", mealHandler.GetMealsDateUser)
}
