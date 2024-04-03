package Postgress

import (
	"context"
	"time"

	conf "github.com/ViPDanger/Golang/Internal/Config"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, config conf.Conf) (pool *pgxpool.Pool, err error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	url := "postgres://" + config.PG_user + ":" + config.PG_password + "@" + config.PG_host + ":" + config.PG_port + "/" + config.PG_bdname
	attempts := config.PG_Con_Attempts

	for attempts > 0 {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, url)

		conf.Err_log(err)
		if err != nil {
			time.Sleep(5 * time.Second)
		}

		attempts--
	}

	conn, err := pgx.Connect(context.Background(), url)

}
