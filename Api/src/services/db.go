package services

import (
	"database/sql"
	"main/src/utils"
)

func GetDb() (*sql.DB, error) {
	// Cadena de conexión a la base de datos
	db, err := sql.Open("mysql", utils.GetUrlMysqlConnection())
	if err != nil {
		utils.ErrorEmail("Connect DataBase", "Error to connect database")
		return nil, err
	}

	// Intentar establecer una conexión
	err = db.Ping()
	if err != nil {
		utils.ErrorEmail("Connect DataBase", "Error to ping database")
		return nil, err
	}

	return db, nil
}
