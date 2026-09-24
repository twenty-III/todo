package db

import (
	"database/sql"
	"fmt"
	"io/fs"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

func OpenSqlite3(dsn string) (*sql.DB, error) {
	return sql.Open("sqlite", dsn)
}

func Migrate(db *sql.DB, fs fs.FS, dir string) error {
	goose.SetBaseFS(fs)
	defer goose.SetBaseFS(nil)
	if err := goose.SetDialect("sqlite3"); err != nil {
		return fmt.Errorf("failed to set dialect: %w", err)
	}
	if err := goose.Up(db, dir); err != nil {
		return fmt.Errorf("failed to goose up: %w", err)
	}
	return nil
}
