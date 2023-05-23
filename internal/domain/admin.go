package domain

import "time"

type Admin struct {
	ID       int
	Email    string
	Name     string
	Password string

	Session
}

type Session struct {
	RefreshToken string
	ExpiresAt    time.Time
}
