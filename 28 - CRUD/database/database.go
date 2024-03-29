package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Mysql's connection driver
)

// Open database connection
func Connect() (*sql.DB, error) {
	connectionString := "root:root@/devbook?charset=utf8&parseTime=true&loc=Local"

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
