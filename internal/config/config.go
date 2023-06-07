package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/subosito/gotenv"
	"golang.org/x/crypto/bcrypt"
)

const (
	defaultServerPort               = "8001"
	defaultServerRWTimeout          = 10 * time.Second
	defaultServerMaxHeaderMegabytes = 1
	defaultAccessTokenTTL           = 15 * time.Minute
	defaultRefreshTokenTTL          = 24 * time.Hour * 30
	defailtCost                     = bcrypt.DefaultCost
)

type Config struct {
	JWT      jwtConfig
	Server   serverConfig
	Database databaseConfig
	Hash     hashConfig
	Redis    RedisConfig
	Kafka    KafkaConfig
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

type hashConfig struct {
	Cost int
}

type RedisConfig struct {
	DB       int
	Host     string
	Port     string
	Password string
}

type KafkaConfig struct {
	Brokers []string
	Topic   string
}

func Init(path string) (*Config, error) {
	err := gotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables from file: %v", err)
	}

	driver := os.Getenv("DRIVER")
	dsn := os.Getenv("DSN_STUDENTS")
	jwtKey := os.Getenv("SIGNIN_KEY")

	host := os.Getenv("HOST_REDIS")
	port := os.Getenv("PORT_REDIS")
	password_redis := os.Getenv("PASSWORD_REDIS")
	db_redis, err := strconv.Atoi(os.Getenv("DB_REDIS"))
	if err != nil {
		return nil, err
	}
	brokerStr := os.Getenv("KAFKA_BROKERS")
	brokers := strings.Split(brokerStr, ",")
	topic := os.Getenv("KAFKA_TOPIC")

	if err != nil {
		return nil, err
	}
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
		Hash: hashConfig{
			Cost: defailtCost,
		},
		Redis: RedisConfig{
			Host:     host,
			Port:     port,
			DB:       db_redis,
			Password: password_redis,
		},
		Kafka: KafkaConfig{
			Brokers: brokers,
			Topic:   topic,
		},
	}, nil

}
