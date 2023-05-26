package domain

import "time"

type Admin struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Name     string `json:"name" binding:"required,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`

	RefreshToken string    `json:"-"`
	ExpiresAt    time.Time `json:"-"`
}

type Session struct {
	RefreshToken string `json:"refreshtoken" binding:"required"`
	ExpiresAt    time.Time
}

type Token struct {
	RefreshToken string `json:"refreshtoken"`
	AccessToken  string `json:"accesstoken"`
}
