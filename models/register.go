package models

import "github.com/celso-alexandre/learning-go-02-rest-api-with-auth/db"

type Register struct {
	Id      int64 `json:"id"`
	EventId int64 `json:"eventId"`
	UserId  int64 `json:"userId"`
}

func (r *Register) CreateRegister() error {
	row, err := db.DB.Exec("insert into registrations (eventId, userId) values (?, ?)", r.EventId, r.UserId)
	if err != nil {
		return err
	}
	r.Id, err = row.LastInsertId()
	if err != nil {
		return err
	}
	return err
}

func (r *Register) DeleteRegister() error {
	_, err := db.DB.Exec("delete from registrations where eventId = ? and userId = ?", r.EventId, r.UserId)
	return err
}

func GetRegistersByEventId(eventId, userId int64) (*[]Register, error) {
	rows, err := db.DB.Query(`
		select id, eventId, userId 
		from registrations 
		where eventId = ?
		  and userId  = ?
	`, eventId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	registers := []Register{}
	for rows.Next() {
		var register Register
		rows.Scan(&register.Id, &register.EventId, &register.UserId)
		registers = append(registers, register)
	}

	return &registers, nil
}

func GetAllRegisters(userId int64) (*[]Register, error) {
	rows, err := db.DB.Query(`
		select id, eventId, userId 
		from registrations 
		where userId  = ?
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	registers := []Register{}
	for rows.Next() {
		var register Register
		rows.Scan(&register.Id, &register.EventId, &register.UserId)
		registers = append(registers, register)
	}

	return &registers, nil
}
