package models

import (
	_ "embed"
	"time"
)

//go:embed meal.sql
var MealTable string

type Meal struct {
	User     string    `json:"user"`
	Name     string    `json:"name"`
	Calroies int       `json:"calories"`
	AddedAt  time.Time `json:"added_at"`
}
