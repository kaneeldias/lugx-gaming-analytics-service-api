package main

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"os"
	"sync"
)

var (
	db   driver.Conn
	once sync.Once
)

func createConnection() (driver.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", os.Getenv("CLICKHOUSE_HOST"), os.Getenv("CLICKHOUSE_PORT"))},
		Auth: clickhouse.Auth{
			Database: os.Getenv("CLICKHOUSE_DATABASE"),
			Password: os.Getenv("CLICKHOUSE_PASSWORD"),
		},
		Protocol: clickhouse.HTTP,
	})
	if err != nil {
		return nil, err
	}
	v, err := conn.ServerVersion()
	if err != nil {
		return nil, fmt.Errorf("failed to get server version: %w", err)
	}
	fmt.Printf("Connected to ClickHouse server version: %s\n", v)

	return conn, nil
}

func GetDatabaseConnection() (driver.Conn, error) {
	var err error
	once.Do(func() {
		db, err = createConnection()
	})

	if err != nil {
		return nil, fmt.Errorf("error creating database connection: %w", err)
	}

	if db == nil {
		return nil, fmt.Errorf("failed to create or retrieve database connection")
	}

	return db, err
}

func SavePageView(path string, ipAddress string) error {
	conn, err := GetDatabaseConnection()
	if err != nil {
		return fmt.Errorf("error getting database connection: %w", err)
	}

	query := "INSERT INTO page_views (path, ip_address) VALUES (?, ?)"
	if err := conn.Exec(context.Background(), query, path, ipAddress); err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}
