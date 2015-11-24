package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type ConnectionDetails struct {
	Host string
	Port string
	Username string
	Password string
	Database string
}

func (conds *ConnectionDetails) connect() {
	var dsn string = conds.Username + ":" + conds.Password + "@tcp(" + conds.Host +
			":" + conds.Port + ")/" + conds.Database
	db, err := sql.Open("mysql", dsn)
}

