package config

type Config struct {
	Database struct {
		Driver string
		DSN    string
	}
}

func NewConfig(path string) (*Config, error) {
	return nil, nil
}
