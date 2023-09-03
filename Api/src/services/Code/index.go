package code

import (
	"fmt"
	"main/src/services"
	user "main/src/services/User"
	"main/src/types"
)

func CreateCode(code types.Code) error {
	db, er := services.GetDb()
	if er != nil {
		return er
	}

	_, r := user.GetUser(code.IdUser)
	if r != nil {
		return r
	}

	_, err := db.Exec("INSERT INTO code (code, order, isplayed, iduser, idlist) VALUES (?, ?, ?, ?, ?)",
		code.Code, code.Order, code.IsPlayed, code.IdUser, code.IdList)
	defer db.Close()
	return err
}

func UpdateCode(id string, code types.Code) error {
	db, er := services.GetDb()
	if er != nil {
		return er
	}
	_, err := db.Exec("UPDATE code SET code=?, order=?, isplayed=?, iduser=?, idlist=? WHERE id=?",
		code.Code, code.Order, code.IsPlayed, code.IdUser, code.IdList)
	defer db.Close()
	return err
}

func DeleteCode(id string) error {
	db, er := services.GetDb()
	if er != nil {
		return er
	}
	_, err := db.Exec("DELETE FROM code WHERE id=?", id)
	defer db.Close()
	return err
}

func GetCode(id string) (types.Code, error) {
	db, er := services.GetDb()
	if er != nil {
		return types.CodeMuckUp, er
	}
	var code types.Code
	err := db.QueryRow("SELECT code, order, isplayed, iduser, idlist FROM code WHERE id=?", id).
		Scan(&code.Id, &code.Code, &code.Order, &code.IsPlayed, &code.IdUser, &code.IdList)
	defer db.Close()
	return code, err
}

func GetCodes(limit string, offset string) (types.Paginate, error) {
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
	query := fmt.Sprintf("SELECT * FROM code limit %s offset %s;", limitt, offsett)
	rows, err := db.Query(query)
	if err != nil {
		return types.Paginate{}, err
	}
	defer rows.Close()

	var codeList []types.Code

	for rows.Next() {
		var code types.Code
		err := rows.Scan(&code.Id, &code.Code, &code.Order, &code.IsPlayed, &code.IdUser, &code.IdList)
		if err != nil {
			return types.Paginate{}, err
		}
		codeList = append(codeList, code)
	}

	return types.Paginate{
		Data:   codeList,
		Limit:  limitt,
		Offset: offsett,
	}, nil
}

func GetCodeByCode(code string) (types.Code, error) {
	db, er := services.GetDb()
	if er != nil {
		return types.CodeMuckUp, er
	}
	var coded types.Code
	err := db.QueryRow("SELECT code, order, isplayed, iduser, idlist FROM code WHERE code=?", code).
		Scan(&coded.Id, &coded.Code, &coded.Order, &coded.IsPlayed, &coded.IdUser, &coded.IdList)
	defer db.Close()
	return coded, err
}
