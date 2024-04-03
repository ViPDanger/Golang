package Postgress

import (
	"database/sql"
	"fmt"

	conf "github.com/ViPDanger/Golang/Internal/Config"
	_ "github.com/lib/pq"
)

func PG_connect(config conf.Conf) *sql.DB {

	psqlconn := "host=" + config.PG_host + " port=" + config.PG_port + " user=" + config.PG_user + " password=" + config.PG_password + " dbname=" + config.PG_bdname + " sslmode=disable"
	// open database
	db, err := sql.Open("postgres", psqlconn)
	conf.Err_log(err)
	defer db.Close()

	// Проверка БД
	err = db.Ping()
	conf.Err_log(err)

	fmt.Println("Connected!")
	return db
}
