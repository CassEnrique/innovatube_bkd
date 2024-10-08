package main

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

const MANDATORY = "the env is mandatory"

func validateEnvs() error {
	if strings.TrimSpace(os.Getenv("SERVER_PORT")) == "" {
		return errors.New("the SERVER_PORT env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("ALLOWED_ORIGINS")) == "" {
		return errors.New("the ALLOWED_ORIGINS env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("ALLOWED_METHODS")) == "" {
		return errors.New("the ALLOWED_METHODS env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("JWT_SECRET_KEY")) == "" {
		return errors.New("the JWT_SECRET_KEY env is mandatory")
	}

	// Database
	if strings.TrimSpace(os.Getenv("DB_USER")) == "" {
		return errors.New(MANDATORY)
	}
	if strings.TrimSpace(os.Getenv("DB_PASSWORD")) == "" {
		return errors.New(MANDATORY)
	}
	if strings.TrimSpace(os.Getenv("DB_HOST")) == "" {
		return errors.New(MANDATORY)
	}
	if strings.TrimSpace(os.Getenv("DB_PORT")) == "" {
		return errors.New(MANDATORY)
	}
	if strings.TrimSpace(os.Getenv("DB_NAME")) == "" {
		return errors.New(MANDATORY)
	}
	if strings.TrimSpace(os.Getenv("DB_SSL_MODE")) == "" {
		return errors.New(MANDATORY)
	}

	return nil
}
