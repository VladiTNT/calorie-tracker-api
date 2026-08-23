package queries

import (
	"context"

	"github.com/VladiTNT/calorie-tracker-api/internal/database/models"
)

func InsertProfile(e Executor, ctx context.Context, p *models.UserProfile) error {
	query := "INSERT INTO profile (user_name, target_calories) VALUES (?, ?);"

	_, err := e.ExecContext(ctx, query, p.Name, p.TargetCalories)
	if err != nil {
		return err
	}

	return nil
}

func GetProfileByName(e Executor, ctx context.Context, userName string) (models.UserProfile, error) {
	query := "SELECT * FROM profile WHERE user_name = ?;"

	var p models.UserProfile
	err := e.QueryRowContext(ctx, query, userName).Scan(&p.Name, &p.TargetCalories)
	if err != nil {
		return p, err
	}

	return p, nil
}
