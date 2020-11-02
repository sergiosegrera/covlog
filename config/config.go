package config

import "os"

type Config struct {
	HttpPort string
	RedisAddress string
	RedisPassword string
	TwilioId string
	TwilioToken string
}

func New() *Config {
	return &Config{
		HttpPort: ParseEnv("HTTP_PORT", "8080"),
		RedisAddress: ParseEnv("REDIS_ADDRESS", "redis:6379"),
		RedisPassword: ParseEnv("REDIS_PASSWORD", ""),
		TwilioId: ParseEnv("TWILIO_ID", ""),
		TwilioToken: ParseEnv("TWILIO_TOKEN", ""),
	}
}

func ParseEnv(key string, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return def
}