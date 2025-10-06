package connection

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // registers "pgx" for database/sql
)

var db *sql.DB

// DB returns the initialized pool.
func DB() *sql.DB {
	if db == nil {
		log.Panic("database not initialized: call connection.MustInit() first")
	}
	return db
}

// MustInit initializes the DB pool or exits.
func MustInit() {
	if err := Init(); err != nil {
		log.Fatalf("❌ database init failed: %v", err)
	}
}

// Init initializes the DB pool (idempotent).
func Init() error {
	if db != nil {
		return nil
	}

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		// fall back to parts
		host := getenv("DB_HOST", "localhost")
		port := getenv("DB_PORT", "5432")
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		name := os.Getenv("DB_NAME")
		ssl := getenv("DB_SSLMODE", "require") // RDS default

		// URL-encode password safely
		pw := url.QueryEscape(pass)

		dsn = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			user, pw, host, port, name, ssl,
		)
	}

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}

	// Pool tuning (tweak as needed)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(45 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	// Add connection timeout parameters to DSN if not using full URL
	if os.Getenv("DB_URL") == "" {
		// Rebuild DSN with timeout parameters
		host := getenv("DB_HOST", "localhost")
		port := getenv("DB_PORT", "5432")
		user := os.Getenv("DB_USER")
		pass := url.QueryEscape(os.Getenv("DB_PASS"))
		name := os.Getenv("DB_NAME")
		ssl := getenv("DB_SSLMODE", "require")

		dsn = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s&connect_timeout=30&statement_timeout=30000",
			user, pass, host, port, name, ssl,
		)

		// Reopen with timeout parameters
		_ = sqlDB.Close()
		sqlDB, err = sql.Open("pgx", dsn)
		if err != nil {
			return fmt.Errorf("reopen with timeouts: %w", err)
		}

		// Reapply pool settings
		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(45 * time.Minute)
		sqlDB.SetConnMaxIdleTime(10 * time.Minute)
	}

	// Fail fast if not reachable
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		_ = sqlDB.Close()
		return fmt.Errorf("ping: %w", err)
	}

	db = sqlDB
	log.Printf("✅ PostgreSQL connected")
	return nil
}

// TestConnection tests the database connection and returns detailed status
func TestConnection() error {
	if db == nil {
		return fmt.Errorf("database not initialized: call connection.MustInit() first")
	}

	// Test basic connectivity
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}

	// Test a simple query
	var result int
	query := "SELECT 1"
	if err := db.QueryRowContext(ctx, query).Scan(&result); err != nil {
		return fmt.Errorf("simple query failed: %w", err)
	}

	if result != 1 {
		return fmt.Errorf("unexpected query result: got %d, expected 1", result)
	}

	// Get connection stats
	stats := db.Stats()
	log.Printf("🔍 Database Connection Stats:")
	log.Printf("   Open Connections: %d", stats.OpenConnections)
	log.Printf("   In Use: %d", stats.InUse)
	log.Printf("   Idle: %d", stats.Idle)
	log.Printf("   Max Open Allowed: %d", stats.MaxOpenConnections)
	log.Printf("   Max Idle Closed: %d", stats.MaxIdleClosed)
	log.Printf("   Total Opened: %d", stats.MaxLifetimeClosed)
	log.Printf("   Total Closed: %d", stats.MaxIdleClosed)

	log.Printf("✅ Database connection test successful")
	return nil
}

// GetConnectionStats returns current database connection statistics
func GetConnectionStats() sql.DBStats {
	if db == nil {
		log.Panic("database not initialized: call connection.MustInit() first")
	}
	return db.Stats()
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
