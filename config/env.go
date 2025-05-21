package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	// PORT returns the server listening port
	PORT string
	// DB returns the name of the sqlite database
	DB string
	// TOKENKEY returns the jwt token secret
	TOKENKEY string
	// TOKENEXP returns the jwt token expiration duration.
	// Should be time.ParseDuration string. Source: https://golang.org/pkg/time/#ParseDuration
	// default: 10h
	TOKENEXP string
)

func Init() {
	_ = godotenv.Load() // Load .env file once, ignore error if not found

	PORT = getEnv("PORT", "5000")
	DB = getEnv("DB", "gotodo.db")
	TOKENKEY = getEnv("TOKEN_KEY", "test")
	TOKENEXP = getEnv("TOKEN_EXP", "10h")
}

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
