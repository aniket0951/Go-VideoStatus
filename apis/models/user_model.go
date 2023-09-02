package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID              uuid.UUID    `json:"id"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Contact         string       `json:"contact"`
	Password        string       `json:"password"`
	UserType        string       `json:"user_type"`
	IsAccountActive sql.NullBool `json:"is_account_active"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
}