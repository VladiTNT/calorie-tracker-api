package queries

import (
	"context"

	"github.com/VladiTNT/calorie-tracker-api/internal/database/models"
)

func InsertUser(e Executor, ctx context.Context, u *models.User) error {
	query := "INSERT INTO user (user_name, user_email, user_pass) VALUES (?, ?, ?);"

	_, err := e.ExecContext(ctx, query, u.Name, u.Email, u.Pass)
	if err != nil {
		return err
	}

	return nil
}

func ValidateUser(e Executor, ctx context.Context, u *models.User) bool {
	query := "SELECT user_pass FROM user WHERE user_name = ?;"

	var pass string
	err := e.QueryRowContext(ctx, query, u.Name).Scan(&pass)
	if err != nil {
		return false
	}

	if u.Pass != pass {
		return false
	}

	return true
}
