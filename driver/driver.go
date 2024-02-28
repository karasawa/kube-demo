package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_unix.go: %s environment variable not set.", k)
		}
		return v
	}
	var (
		dbUser         = mustGetenv("DB_USER")
		dbPwd          = mustGetenv("DB_PASSWORD")
		dbName         = mustGetenv("DB_NAME")
		// unixSocketPath = mustGetenv("INSTANCE_UNIX_SOCKET") // e.g. '/cloudsql/project:region:instance'
	)
	// dbURI := fmt.Sprintf("%s:%s@unix(%s)/%s?parseTime=true",
	// 	dbUser, dbPwd, unixSocketPath, dbName)
    dbURI := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPwd, dbName)
	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("sql.Open: %s", err)
	}
	return dbPool, nil
}
