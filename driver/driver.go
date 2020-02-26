package driver

import (
	"database/sql"
	"github.com/lib/pq"
	"os"
	"log"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	// get pgUrl
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANT_SQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)
	// db ping
	err = db.Ping()
	logFatal(err)

	return db
}

