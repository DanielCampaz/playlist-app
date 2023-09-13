package list

import (
	"fmt"
	"main/src/services"
	user "main/src/services/User"
	"main/src/types"
)

func CreateList(list types.List) error {
	db, er := services.GetDb()
	if er != nil {
		return er
	}

	_, r := user.GetUser(list.IdUser)
	if r != nil {
		return r
	}

	_, err := db.Exec("INSERT INTO list (name, iduser, act, counts) VALUES (?, ?, ?, ?)",
		list.Name, list.IdUser, list.Act, list.Counts)
	defer db.Close()
	return err
}

func UpdateList(id string, list types.List) error {
	db, er := services.GetDb()
	if er != nil {
		return er
	}
	_, err := db.Exec("UPDATE list SET name=?, iduser=?, act=?, counts=? WHERE id=?",
		list.Name, list.IdUser, list.Act, list.Counts, id)
	defer db.Close()
	return err
}

func UpdateCount(id string) error {
	list, erro := GetList(id)
	if erro != nil {
		return erro
	}
	db, er := services.GetDb()
	if er != nil {
		return er
	}
	conut := list.Counts + 1
	_, err := db.Exec("UPDATE list SET name=?, iduser=?, act=?, counts=? WHERE id=?",
		list.Name, list.IdUser, list.Act, conut, id)
	defer db.Close()
	return err
}

func UpdateAct(id string) error {
	list, erro := GetList(id)
	if erro != nil {
		return erro
	}
	db, er := services.GetDb()
	if er != nil {
		return er
	}
	act := list.Act + 1
	_, err := db.Exec("UPDATE list SET name=?, iduser=?, act=?, counts=? WHERE id=?",
		list.Name, list.IdUser, act, list.Counts, id)
	defer db.Close()
	return err
}

func DeleteList(id string) error {
	db, er := services.GetDb()
	if er != nil {
		return er
	}
	_, err := db.Exec("DELETE FROM list WHERE id=?", id)
	defer db.Close()
	return err
}

func GetList(id string) (types.List, error) {
	db, er := services.GetDb()
	if er != nil {
		return types.ListMuckUp, er
	}
	var list types.List
	err := db.QueryRow("SELECT id, name, iduser, act, counts FROM list WHERE id=?", id).
		Scan(&list.Id, &list.Name, &list.IdUser, &list.Act, &list.Counts)
	defer db.Close()
	return list, err
}

func GetLists(limit string, offset string) (types.Paginate, error) {
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
	query := fmt.Sprintf("SELECT * FROM list limit %s offset %s;", limitt, offsett)
	rows, err := db.Query(query)
	if err != nil {
		return types.Paginate{}, err
	}
	defer rows.Close()

	var listLis []types.List

	for rows.Next() {
		var list types.List
		err := rows.Scan(&list.Id, &list.Name, &list.IdUser, &list.Act, &list.Counts)
		if err != nil {
			return types.Paginate{}, err
		}
		listLis = append(listLis, list)
	}

	return types.Paginate{
		Data:   listLis,
		Limit:  limitt,
		Offset: offsett,
	}, nil
}

func GetListByName(name string) (types.List, error) {
	db, er := services.GetDb()
	if er != nil {
		return types.ListMuckUp, er
	}
	var list types.List
	err := db.QueryRow("SELECT id, name, iduser, act, counts FROM list WHERE name=?", name).
		Scan(&list.Id, &list.Name, &list.IdUser, &list.Act, &list.Counts)
	defer db.Close()
	return list, err
}
