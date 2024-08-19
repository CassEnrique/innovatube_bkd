package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

func newDBConnection() (*gorm.DB, error) {
	min := 3
	max := 100

	minConn := os.Getenv("DB_MIN_CONN")
	maxConn := os.Getenv("DB_MAX_CONN")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	if minConn != "" {
		v, err := strconv.Atoi(minConn)
		if err != nil {
			log.Println("Warning: DB_MIN_CONN has not a valid value, we will set min connections to", min)
		} else {
			if v >= min && v <= max {
				min = v
			}
		}
	}
	if maxConn != "" {
		v, err := strconv.Atoi(maxConn)
		if err != nil {
			log.Println("Warning: DB_MAX_CONN has not a valid value, we will set max connections to", max)
		} else {
			if v >= min && v <= max {
				max = v
			}
		}
	}

	connString := makeURL(user, pass, host, port, dbName, sslMode)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s %w", "gorm.Open()", err)
	}

	return db, nil
}

func makeURL(user, pass, host, port, dbName, sslMode string) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s TimeZone=America/Mexico_City",
		user,
		pass,
		host,
		port,
		dbName,
		sslMode,
	)
}
