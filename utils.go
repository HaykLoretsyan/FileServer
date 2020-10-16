package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)


func init() {

	// Get current directory
	dir, err := os.Getwd()
	Must(err)

	// Load .env file from project root directory
	Must(godotenv.Load( filepath.Join(dir, ".env")))
}

func GetEnvVar(name string) string {

	// Return environment variable with name 'name'
	return os.Getenv(name)
}

func Must(err error) {

	if err != nil {
		// Log error and exit
		log.Fatal(err)
	}
}
