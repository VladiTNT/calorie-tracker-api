package app

import "github.com/VladiTNT/calorie-tracker-api/internal/handlers"

func (ctapi *CalorieTrackerAPI) RegisterRoutes() {
	ctapi.Router.HandleFunc("GET /", handlers.TestHandler)
}
