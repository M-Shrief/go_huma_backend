package config

import (
	"log"

	"github.com/joho/godotenv"
)

var envs map[string]string

func LoadENV() {
	var err error
	envs, err = godotenv.Read(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	assignValues()
}

// General
var (
	HOST string
	PORT string
)

func assignValues() {
	// General
	HOST = envs["HOST"]
	PORT = envs["PORT"]
}
