// internal/storage/database.go
package internal

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabaseConnection() *sql.DB {
	// dsn: Data Source Name
	dsn := "root:@tcp(127.0.0.1:3306)/user_management"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
