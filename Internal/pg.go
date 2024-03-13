package Internal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func PG_connect() *sql.DB {
	config := Read_Config()
	// connection string
	psqlconn := "host=" + config.PG_host + " port=" + config.PG_port + " user=" + config.PG_user + " password=" + config.PG_password + " dbname=" + config.PG_bdname + " sslmode=disable"
	// open database
	db, err := sql.Open("postgres", psqlconn)
	err_log(err)
	defer db.Close()

	// Проверка БД
	err = db.Ping()
	err_log(err)

	fmt.Println("Connected!")
	return db
}
