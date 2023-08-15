package user

import (
	"fmt"
	"main/src/services"
	"main/src/types"
)

func CreateUser(u types.User) error {
	db, er := services.GetDb()
	if er != nil {
		return er
	}
	_, err := db.Exec("INSERT INTO users (name, lastname, email, password) VALUES (?, ?, ?, ?)",
		u.Name, u.LastName, u.Email, u.Password)
	defer db.Close()
	return err
}

func UpdateUser(id string, u types.User) error {
	db, er := services.GetDb()
	if er != nil {
		return er
	}
	_, err := db.Exec("UPDATE users SET name=?, lastname=?, email=?, password=? WHERE id=?",
		u.Name, u.LastName, u.Email, u.Password, id)
	defer db.Close()
	return err
}

func DeleteUser(id string) error {
	db, er := services.GetDb()
	if er != nil {
		return er
	}
	_, err := db.Exec("DELETE FROM users WHERE id=?", id)
	defer db.Close()
	return err
}

func GetUser(id string) (types.User, error) {
	db, er := services.GetDb()
	if er != nil {
		return types.UserMuckUp, er
	}
	var user types.User
	err := db.QueryRow("SELECT id, name, lastname, email, password FROM users WHERE id=?", id).
		Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.Password)
	defer db.Close()
	return user, err
}

func GetUsers(limit string, offset string) (types.Paginate, error) {
	db, er := services.GetDb()
	var limitt = "10"
	var offsett = "0"
	if er != nil {
		return types.Paginate{}, er
	}
	if limit != "" {
		limitt = limit
	}
	if offset != "" {
		offsett = offset
	}
	query := fmt.Sprintf("SELECT * FROM users limit %s offset %s;", limitt, offsett)
	rows, err := db.Query(query)
	if err != nil {
		return types.Paginate{}, err
	}
	defer rows.Close()

	var userList []types.User

	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			return types.Paginate{}, err
		}
		userList = append(userList, user)
	}

	return types.Paginate{
		Data:   userList,
		Limit:  limitt,
		Offset: offsett,
	}, nil
}
