package database

import (
	"database/sql"
	"fmt"

	"go.elastic.co/apm/module/apmsql/v2"
	_ "go.elastic.co/apm/module/apmsql/v2/pq"
)

// with APM
func NewDBClient(host, userName, password, sslMode, port, databaseName string) (*sql.DB, error) {
	connString := getConnectionString(host, userName, password, sslMode, port, databaseName)
	db, err := apmsql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getConnectionString(host, userName, password, sslMode, port, databaseName string) string {
	connectionStringTemplate := "host=%s port=%s sslmode=%s user=%s password='%s' dbname=%s "
	return fmt.Sprintf(
		connectionStringTemplate,
		host,
		port,
		sslMode,
		userName,
		password,
		databaseName)
}
