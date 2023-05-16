package main

import (
	"log"

	"github.com/begenov/test-task-backend/students-app/internal/config"
)

const (
	path_config = ""
)

func main() {
	cfg, err := config.NewConfig(path_config)
	if err != nil {
		log.Fatalf("can't load config: %v", err)
		return
	}

}
