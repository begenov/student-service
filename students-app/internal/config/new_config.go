package config

import (
	"fmt"
	"time"

	"github.com/subosito/gotenv"
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
}

func NewConfig(path string) (*Config, error) {
	err := gotenv.Load(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables from file: %v", err)
	}
	return nil, nil
}
