package types

import (
	"main/src/services"
)

type User struct {
	Id       int16  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) CreateUser() error {
	db, er := services.GetDb()
	if er != nil {
		return er
	}
	_, err := db.Exec("INSERT INTO users (name, lastname, email, password) VALUES (?, ?, ?, ?)",
		u.Name, u.LastName, u.Email, u.Password)
	defer db.Close()
	return err
}

func (u User) updateUser() error {
	db, er := services.GetDb()
	if er != nil {
		return er
	}
	_, err := db.Exec("UPDATE users SET name=?, lastname=?, email=?, password=? WHERE email=?",
		u.Name, u.LastName, u.Email, u.Password, u.Email)
	return err
}

func (u User) deleteUser() error {
	db, er := services.GetDb()
	if er != nil {
		return er
	}
	_, err := db.Exec("DELETE FROM users WHERE email=?", u.Email)
	return err
}

func (u User) getUser() (User, error) {
	db, er := services.GetDb()
	if er != nil {
		return u, er
	}
	var user User
	err := db.QueryRow("SELECT id, name, lastname, email, password FROM users WHERE email=?", u.Email).
		Scan(&user.Name, &user.LastName, &user.Email, &user.Password)
	return user, err
}
