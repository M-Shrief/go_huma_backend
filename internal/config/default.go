package config

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
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
	APP_ENV string
	HOST    string
	PORT    string
)

// JWT
var (
	JWT_PRIVATE *rsa.PrivateKey
	JWT_PUBLIC  *rsa.PublicKey
)

// Database
var (
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
)

func assignValues() {
	// General
	APP_ENV = envs["APP_ENV"]
	HOST = envs["HOST"]
	PORT = envs["PORT"]

	// Database
	DB_HOST = envs["DB_HOST"]
	DB_PORT = envs["DB_PORT"]
	DB_USER = envs["DB_USER"]
	DB_NAME = envs["DB_NAME"]
	DB_PASSWORD = envs["DB_PASSWORD"]

	// JWT
	assignJWTKeys()
}

func assignJWTKeys() {
	// JWT
	jwt_Private_File, err := os.ReadFile("./jwtRSA256-private.pem")
	if err != nil {
		fmt.Printf("couldn't parse private JWT secret file: %v\n", err)
		os.Exit(1)
	}

	JWT_PRIVATE, err = jwt.ParseRSAPrivateKeyFromPEM(jwt_Private_File)
	if err != nil {
		fmt.Printf("couldn't parse private JWT secret: %v\n", err)
		os.Exit(1)
	}

	// JWT
	jwt_Public_File, err := os.ReadFile("./jwtRSA256-public.pem")
	if err != nil {
		fmt.Printf("couldn't read public JWT secret file: %v\n", err)
		os.Exit(1)
	}

	JWT_PUBLIC, err = jwt.ParseRSAPublicKeyFromPEM(jwt_Public_File)
	if err != nil {
		fmt.Printf("couldn't parse public JWT secret: %v\n", err)
		os.Exit(1)
	}
}
