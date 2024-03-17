package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Value     string
	Type      string
	ModelID   int
	ModelType int
	ExpiresAt time.Time
}

func (t Token) HasExpired() bool {
	return t.ExpiresAt.Before(time.Now())
}

const (
	TokenUserActivation string = "user_activation"
	TokenPasswordReset  string = "password_reset"
)
