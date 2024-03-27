package models

import (
	"time"

	"github.com/celso-alexandre/learning-go-02-rest-api-with-auth/db"
)

type Event struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	UserId      int       `json:"userId"`
}

func (e *Event) Create() error {
	const sql = `
		INSERT INTO events (name, description, location, date, userId) 
		VALUES ($1, $2, $3, $4, $5) 
		--RETURNING id
	`
	// row := db.DB.QueryRow(sql, e.Name, e.Description, e.Location, e.Date, e.UserId)
	// return row.Scan(&e.Id)

	// Yup, in this case we only call this same statemente 1 time before .Close()
	// So, there was no real need for Prepare()
	stmt, err := db.DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Date, e.UserId)
	if err != nil {
		return err
	}
	e.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

const dateTimeFormat = "2006-01-02 15:04:05-03:00"

func formatDbDateTime(t string) (time.Time, error) {
	return time.Parse(dateTimeFormat, t)
}

func FindAllEvents() ([]Event, error) {
	const sql = `SELECT id, name, description, location, date, userId FROM events`
	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		var dateString string
		err := rows.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &dateString, &e.UserId)
		if err != nil {
			return nil, err
		}
		e.Date, err = formatDbDateTime(dateString)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func FindEventById(id string) (*Event, error) {
	const sql = `SELECT id, name, description, location, date, userId FROM events WHERE id = $1`
	row := db.DB.QueryRow(sql, id)

	var e Event
	var dateString string
	err := row.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &dateString, &e.UserId)
	if err != nil {
		return nil, err
	}
	e.Date, err = formatDbDateTime(dateString)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
