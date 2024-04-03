package Postgress

import (
	"context"
	"strconv"
	"time"

	"github.com/ViPDanger/Golang/Internal/config"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, conf config.Conf) (pool *pgxpool.Pool, err error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	url := "postgres://" + conf.PG_user + ":" + conf.PG_password + "@" + conf.PG_host + ":" + conf.PG_port + "/" + conf.PG_bdname
	attempts, _ := strconv.Atoi(conf.PG_Con_Attempts)

	for attempts > 0 {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, url)

		config.Err_log(err)
		if err != nil {
			time.Sleep(5 * time.Second)
		}

		attempts--
	}

	return pool, nil

}
