package config

import (
	"fmt"
	"os"
	"time"

	"github.com/subosito/gotenv"
)

const (
	defaultServerPort               = "8000"
	defaultServerRWTimeout          = 10 * time.Second
	defaultServerMaxHeaderMegabytes = 1
	defaultAccessTokenTTL           = 15 * time.Minute
	defaultRefreshTokenTTL          = 24 * time.Hour * 30
)

type Config struct {
	Server struct {
		Port               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		MaxHeaderMegabytes int
	}

	Database struct {
		Driver string
		DSN    string
	}

	JWT struct {
		AccessTokenTTL  time.Duration
		RefreshTokenTTL time.Duration
		SigningKey      string
	}
}

func NewConfig(path string) (*Config, error) {
	err := gotenv.Load(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables from file: %v", err)
	}

	driver := os.Getenv("DRIVER")
	dsn := os.Getenv("DSN_STUDENTS")
	jwtKey := os.Getenv("SIGNIN_KEY")

	return &Config{
		Server: struct {
			Port               string
			ReadTimeout        time.Duration
			WriteTimeout       time.Duration
			MaxHeaderMegabytes int
		}{
			Port:               defaultServerPort,
			ReadTimeout:        defaultServerRWTimeout,
			WriteTimeout:       defaultServerRWTimeout,
			MaxHeaderMegabytes: defaultServerMaxHeaderMegabytes,
		},
		Database: struct {
			Driver string
			DSN    string
		}{
			Driver: driver,
			DSN:    dsn,
		},
		JWT: struct {
			AccessTokenTTL  time.Duration
			RefreshTokenTTL time.Duration
			SigningKey      string
		}{
			AccessTokenTTL:  defaultAccessTokenTTL,
			RefreshTokenTTL: defaultRefreshTokenTTL,
			SigningKey:      jwtKey,
		},
	}, nil
}
