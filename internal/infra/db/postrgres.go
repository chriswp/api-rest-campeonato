package db

import (
	"context"
	"database/sql"
	"fmt"
	config "github.com/chriswp/api-rest-campeonato/configs"
	"github.com/chriswp/api-rest-campeonato/pkg/logger"
	_ "github.com/lib/pq"
	"time"
)

func NewPostgresConnection(ctx context.Context) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBUser,
		config.Envs.DBPass,
		config.Envs.DBDatabase,
	)

	db, err := sql.Open(config.Envs.DBDriver, dsn)
	if err != nil {
		logger.Error("Error opening database connection", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		logger.Error("Error pinging PostgreSQL database", err)
		db.Close()
		return nil, err
	}

	return db, nil
}
