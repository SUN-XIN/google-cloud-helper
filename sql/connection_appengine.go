// +build appengine

package sql

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"google.golang.org/appengine/log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db             *sql.DB
	connectionName = os.Getenv("APPENGINESQL_CONNECTION_NAME")
	user           = os.Getenv("APPENGINESQL_USER")
	password       = os.Getenv("APPENGINESQL_PASSWORD")
	schemaName     = os.Getenv("APPENGINESQL_SCHEMA")
)

func GetConnection(ctx context.Context) (*sql.DB, error) {
	log.Debugf(ctx, "db Appengine")
	return sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/%s?multiStatements=true&parseTime=true",
		user, password, connectionName, schemaName))
}
