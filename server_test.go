package main

import (
	"knowledge/api/database"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestConnectionStringIsSet(t *testing.T) {
	godotenv.Load()
	connectionString := os.Getenv("CONNECTION_STRING")
	if connectionString == "" {
		t.Error("CONNECTION_STRING environment variable not set")
	}
}

func TestCouldConnectToDatabase(t *testing.T) {
	godotenv.Load()
	connectionString := os.Getenv("CONNECTION_STRING")
	if connectionString == "" {
		t.Error("CONNECTION_STRING environment variable not set")
	}

	_, err := database.ConnectDatabase(connectionString)
	if err != nil {
		t.Error(err)
	}
}
