package database

import (
	"database/sql"

	"github.com/VladiTNT/calorie-tracker-api/internal/database/models"
	_ "modernc.org/sqlite"
)

func New(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", databaseUrl)
	if err != nil {
		return nil, err
	}

	// Migrations
	for _, table := range []string{
		models.UserTable,
		models.ProfileTable,
	} {
		_, err := db.Exec(table)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
