package queries

import (
	"context"

	"github.com/VladiTNT/calorie-tracker-api/internal/database/models"
)

func InsertProfile(e Executor, ctx context.Context, p *models.UserProfile) error {
	query := "INSERT INTO user_profile (user_name, target_calories) VALUES (?, ?);"

	_, err := e.ExecContext(ctx, query, p.Name, p.TargetCalories)
	if err != nil {
		return err
	}

	return nil
}

func GetProfileByName(e Executor, ctx context.Context, userName string) (models.UserProfile, error) {
	query := "SELECT * FROM user_profile WHERE user_name = ?;"

	var p models.UserProfile
	err := e.QueryRowContext(ctx, query, userName).Scan(&p.Name, &p.TargetCalories)
	if err != nil {
		return p, err
	}

	return p, nil
}

func UpdateProfileCalories(e Executor, ctx context.Context, userName string, calories int) error {
	query := "UPDATE user_profile SET target_calories = ? WHERE user_name = ?;"

	_, err := e.ExecContext(ctx, query, calories, userName)
	if err != nil {
		return err
	}

	return nil
}
