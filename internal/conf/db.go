package conf

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

func NewDB() (*sql.DB, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		return nil, fmt.Errorf("DB_DRIVER environment variable must be set")
	}

	dsn, err := buildDSN(dbDriver)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	log.Println("✅ Database connection established successfully")
	return db, nil
}

func buildDSN(driver string) (string, error) {
	switch driver {
	case "postgres", "pgsql":
		return buildPostgresDSN()
	case "mysql":
		return buildMySQLDSN()
	case "sqlite", "sqlite3":
		return buildSQLiteDSN()
	default:
		return "", fmt.Errorf("unsupported database driver: %s. Supported: postgres, mysql, sqlite", driver)
	}
}

func buildPostgresDSN() (string, error) {
	required := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	if err := validateEnvVars(required); err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		getSSLMode(),
	), nil
}

func buildMySQLDSN() (string, error) {
	required := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	if err := validateEnvVars(required); err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		getTLSConfig(),
	), nil
}

func buildSQLiteDSN() (string, error) {
	required := []string{"DB_NAME"}
	if err := validateEnvVars(required); err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"file:%s.db?_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)",
		os.Getenv("DB_NAME"),
	), nil
}

func validateEnvVars(keys []string) error {
	for _, key := range keys {
		if os.Getenv(key) == "" {
			return fmt.Errorf("required environment variable '%s' is not set", key)
		}
	}
	return nil
}

func getSSLMode() string {
	if mode := os.Getenv("DB_SSL_MODE"); mode != "" {
		return mode
	}
	return "disable"
}

func getTLSConfig() string {
	if tls := os.Getenv("DB_TLS"); tls != "" {
		return tls
	}
	return "preferred"
}
