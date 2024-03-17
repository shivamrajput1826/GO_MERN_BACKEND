package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Identifier string
	UserId     uint
	ExpiresAt  time.Time
}

func (s Session) HasExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}
