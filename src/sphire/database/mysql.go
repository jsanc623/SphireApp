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

func (connection_details *ConnectionDetails)connect() {
	
}
