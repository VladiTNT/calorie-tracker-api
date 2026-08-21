package models

import _ "embed"

//go:embed user.sql
var UserTable string

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}
