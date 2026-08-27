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

func GetMealsOnDayByUsername(e Executor, ctx context.Context, userName string, day time.Time) ([]models.Meal, error) {
	query := "SELECT * FROM meal WHERE user_name = ?;"

	rows, err := e.QueryContext(ctx, query, userName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meals []models.Meal
	for rows.Next() {
		var m models.Meal
		if err := rows.Scan(&m.User, &m.Name, &m.Calroies, &m.AddedAt); err != nil {
			return nil, err
		}
		meals = append(meals, m)
	}

	var todayMeals []models.Meal

	for _, m := range meals {
		if m.AddedAt.After(day.Truncate(time.Hour*24)) &&
			m.AddedAt.Before(day.Truncate(time.Hour*24).Add(time.Hour*24)) {
			todayMeals = append(todayMeals, m)
		}
	}

	return todayMeals, nil
}
