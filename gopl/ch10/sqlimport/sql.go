package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // enable support for MySQL
	_ "github.com/lib/pq"              // enable support for Postgres
)

func main() {
	db, err := sql.Open("postgres", "test") // OK
	db, err = sql.Open("mysql", "test")     // OK
	db, err = sql.Open("sqlite3", "test")   // returns error: unknown driver "sqlite3"

	fmt.Print(err)
	fmt.Print(db)
}
