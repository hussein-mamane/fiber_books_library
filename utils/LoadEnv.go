package utils

import "github.com/joho/godotenv"

// LoadEnv not used in this project
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
