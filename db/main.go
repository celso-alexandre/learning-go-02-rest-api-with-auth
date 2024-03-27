package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "file:./api.db?_foreign_keys=on")
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
		create table if not exists users (
			id integer not null primary key autoincrement,
			email text not null unique,
			password text not null
		);
		create table if not exists events (
			id integer not null primary key autoincrement,
			name text not null, 
			description text not null, 
			location text not null, 
			date text not null, 
			userId integer not null,

			FOREIGN KEY(userId) REFERENCES users(id)
		);
		create table if not exists registrations (
			id integer not null primary key autoincrement,
			eventId integer not null,
			userId integer not null,

			FOREIGN KEY(eventId) REFERENCES events(id),
			FOREIGN KEY(userId) REFERENCES users(id)
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
