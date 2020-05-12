package main

import (
	"database/sql"
	"fmt"

	// registers PostgreSQL driver
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres" // you can change this to your custom user
	password = "your-password"
	dbname   = "you-database-name" // you can create this manually with PGAdmin
)

// this holds the database connection
// you can use this pointer to the DB
// to create, update, delete, read (CRUD)
// from postgres
var db *sql.DB

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// "defer" ensures this statement is
	// executes "last" when ever this
	// function is invoked, this in turn closes the
	// database connection, you must remove this
	// statement to b able to use the "db" pointer
	// defined as a package level variable.
	defer db.Close()

	// verify connection is still open
	// reconnects if it isn't
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return db
}
