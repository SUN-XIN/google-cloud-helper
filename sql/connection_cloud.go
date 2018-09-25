// +build !appengine

package sql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db         *sql.DB
	addr       = os.Getenv("CLOUDSQL_CONNECTION_ADDR")
	user       = os.Getenv("CLOUDSQL_USER")
	password   = os.Getenv("CLOUDSQL_PASSWORD")
	schemaName = os.Getenv("CLOUDSQL_SCHEMA")
)

func GetConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true&parseTime=true",
		user,
		password,
		addr,
		schemaName)
	return sql.Open("mysql", dsn)
}
