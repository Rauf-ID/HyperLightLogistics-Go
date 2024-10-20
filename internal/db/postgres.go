package db

import (
	"HyperLightLogistics-Go/internal/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresDB struct {
	Conn *pgxpool.Pool
}

func NewPostgresDB(cfg *config.Config) (*PostgresDB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	conn, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return &PostgresDB{Conn: conn}, nil
}

func (db *PostgresDB) Close() {
	db.Conn.Close()
}
