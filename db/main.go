package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")
	if err != nil {
		fmt.Println(err)
		panic("Error opening database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	runMigrations()
}

func runMigrations() {
	sqlStmt := `
		create table if not exists events (
			id integer not null primary key autoincrement,
			name text not null, 
			description text not null, 
			location text not null, 
			date text not null, 
			userId integer not null
		)
	`
	// if DB == nil {
	// 	fmt.Println("DB is nil")
	// 	return
	// }
	_, err := DB.Exec(sqlStmt)
	if err != nil {
		fmt.Println(err)
		panic("Error running migrations")
	}
}
