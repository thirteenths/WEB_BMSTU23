package postgres

import (
	"context"
	"github.com/jackc/pgx"
	"time"
)

type Connect struct {
	db *pgx.ConnPool
	tx *pgx.Tx
}

func NewConnect(url string) (*Connect, error) {
	cfg, err := pgx.ParseConnectionString(url)

	cpfg := pgx.ConnPoolConfig{ConnConfig: cfg, MaxConnections: 20, AcquireTimeout: time.Second * 30}
	cxt := context.Background()
	db, err := CreateConnect(cxt, cpfg)

	if err != nil {
		return &Connect{}, err
	}
	return &Connect{db: db}, err
}

func CreateConnect(ctx context.Context, cfg pgx.ConnPoolConfig) (*pgx.ConnPool, error) {
	return pgx.NewConnPool(cfg)
}
