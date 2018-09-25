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
	connectionName = os.Getenv("CLOUDSQL_CONNECTION_NAME")
	user           = os.Getenv("CLOUDSQL_USER")
	password       = os.Getenv("CLOUDSQL_PASSWORD")
	schemaName     = os.Getenv("CLOUDSQL_SCHEMA")
)

func GetConnection(ctx context.Context) (*sql.DB, error) {
	log.Debugf(ctx, "db Appengine")
	return sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/%s?multiStatements=true&parseTime=true",
		user, password, connectionName, schemaName))
}
