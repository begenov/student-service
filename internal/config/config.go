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
	JWT      jwtConfig
	Server   serverConfig
	Database databaseConfig
}

type jwtConfig struct {
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	SigningKey      string
}

type serverConfig struct {
	Port               string
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxHeaderMegabytes int
}

type databaseConfig struct {
	Driver string
	DSN    string
}

func Init(path string) (*Config, error) {
	err := gotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables from file: %v", err)
	}

	driver := os.Getenv("DRIVER")
	dsn := os.Getenv("DSN_STUDENTS")
	jwtKey := os.Getenv("SIGNIN_KEY")

	return &Config{
		JWT: jwtConfig{
			AccessTokenTTL:  defaultAccessTokenTTL,
			RefreshTokenTTL: defaultRefreshTokenTTL,
			SigningKey:      jwtKey,
		},
		Server: serverConfig{
			Port:               defaultServerPort,
			ReadTimeout:        defaultServerRWTimeout,
			WriteTimeout:       defaultServerRWTimeout,
			MaxHeaderMegabytes: defaultServerMaxHeaderMegabytes,
		},
		Database: databaseConfig{
			Driver: driver,
			DSN:    dsn,
		},
	}, nil

}
