package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"time"
)

type Connect struct {
	db *pgx.ConnPool
	tx *pgx.Tx
}

func NewConnect(user, password, database string) (*Connect, error) {
	//cfg := pgx.ConnConfig{User: "postgres", Password: "password", Database: "postgres", Port: 5438}
	//cfg := pgx.ConnConfig{User: user, Password: password, Database: database, Port: 5432}

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"standup",
		"password",
		"postgres",
		5432,
		"postgres")

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
