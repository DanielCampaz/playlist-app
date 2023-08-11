package main

import (
	"database/sql"
	"fmt"
	usercontroller "main/src/controllers/UserController"
	"main/src/envirimoents"
	"main/src/utils"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	// Cadena de conexión a la base de datos
	db, err := sql.Open("mysql", utils.GetUrlMysqlConnection())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Intentar establecer una conexión
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successful connection to MySQL database")
	// http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	response := Response{Message: "¡Hola desde la API de Go!"}

	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)

	// 	json.NewEncoder(w).Encode(response)
	// })

	for _, value := range usercontroller.UC {
		http.HandleFunc(value.Url, value.Control)
	}

	fmt.Println("Api Listen in port 8080")
	http.ListenAndServe(envirimoents.GetPort(), nil)

}
