package queries

import (
	"context"
	"time"

	"github.com/VladiTNT/calorie-tracker-api/internal/database/models"
)

func InsertMeal(e Executor, ctx context.Context, m *models.Meal) error {
	query := "INSERT INTO meal (user_name, meal_name, meal_calories) VALUES (?, ?, ?);"

	_, err := e.ExecContext(ctx, query, m.User, m.Name, m.Calroies)
	if err != nil {
		return err
	}

	return nil
}

func GetMealsByUserByDay(e Executor, ctx context.Context, userName string, day time.Time) ([]models.Meal, error) {
	_ = "SELECT * FROM meal WHERE user_name = ? AND added_at"
	return []models.Meal{}, nil
}
