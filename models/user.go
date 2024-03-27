package models

import "github.com/celso-alexandre/learning-go-02-rest-api-with-auth/db"

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Create() error {
	const sql = `
		INSERT INTO users (email, password) 
		VALUES ($1, $2)
	`
	stmt, err := db.DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	u.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}
	u.Password = ""
	return nil
}

func FindUserByEmail(email string) (*User, error) {
	const sql = `
		SELECT id, email, password
		FROM users
		WHERE email = $1
	`
	row := db.DB.QueryRow(sql, email)
	var u User
	err := row.Scan(&u.Id, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func FindUserById(id int64) (*User, error) {
	const sql = `
		SELECT id, email, password
		FROM users
		WHERE id = $1
	`
	row := db.DB.QueryRow(sql, id)
	var u User
	err := row.Scan(&u.Id, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = ""
	return &u, nil
}

func (u *User) Update() error {
	const sql = `
		UPDATE users 
		SET email = $1, password = $2
		WHERE id = $3
	`
	_, err := db.DB.Exec(sql, u.Email, u.Password, u.Id)
	u.Password = ""
	return err
}

func FindAllUsers() ([]User, error) {
	const sql = `SELECT id, email, password FROM users`
	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Email, &u.Password)
		if err != nil {
			return nil, err
		}
		u.Password = ""
		users = append(users, u)
	}
	return users, nil
}

func DeleteUser(id int64) error {
	const sql = `DELETE FROM users WHERE id = $1`
	_, err := db.DB.Exec(sql, id)
	return err
}
