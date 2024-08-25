package config

type config struct {
	Environment  uint8    `yaml:"environment" validate:"gte=0,lte=2"`
	LogLevel     uint32   `yaml:"log_level" validate:"required"`
	DBUrl        string   `yaml:"db_url" validate:"required"`
	DBName       string   `yaml:"db_name" validate:"required"`
	Cors         []string `yaml:"cors" validate:"required"`
	Address      string   `yaml:"address" validate:"required"`
	Token        string   `yaml:"token" validate:"required"`
	GoogleAPIKey string   `yaml:"google_api_key" validate:"required"`
}
