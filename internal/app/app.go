package app

import (
	"database/sql"
	"fmt"
	"learn/internal/config"
	"learn/internal/db"
	"learn/internal/migrations"
	"learn/internal/todo"
	"net/http"
)

const pref = "/api/v1"

type App struct {
	DB  *sql.DB
	Cfg *config.Config
}

func New(cfg *config.Config) (*App, error) {
	sqlite3Db, err := db.OpenSqlite3(cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := db.Migrate(sqlite3Db, migrations.FS, "."); err != nil {
		return nil, fmt.Errorf("failed to migrate DB: %w", err)
	}

	return &App{
		DB:  sqlite3Db,
		Cfg: cfg,
	}, nil
}

func (a *App) Run() error {
	mux := http.NewServeMux()

	todo.RegisterRoutes(mux, pref, a.DB)

	server := http.Server{
		Addr:    ":" + a.Cfg.Port,
		Handler: mux,
	}

	return server.ListenAndServe()
}
