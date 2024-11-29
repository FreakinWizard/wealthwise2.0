package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Address  string
	Database string
	User     string
	Password string
}

func InitDB(cfg *DBConfig) (*sql.DB, error) {
	dsn := cfg.formatDSN()
	log.Printf("Database DSN: '%s'\n", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (cfg *DBConfig) formatDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		cfg.User,
		cfg.Password,
		cfg.Address,
		cfg.Database,
	)
}
