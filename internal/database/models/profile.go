package models

import _ "embed"

//go:embed profile.sql
var ProfileTable string

type UserProfile struct {
	Name           string `json:"name"`
	TargetCalories int    `json:"target_calories"`
}
