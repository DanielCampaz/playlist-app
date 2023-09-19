package main

import (
	"database/sql"
	"fmt"
	"log"
	authcontroller "main/src/controllers/AuthController"
	listcontroller "main/src/controllers/ListController"
	usercontroller "main/src/controllers/UserController"
	"main/src/envirimoents"
	"main/src/utils"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

// type Response struct {
// 	Message string `json:"message"`
// }

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

	routes := mux.NewRouter()

	// AUTH
	for _, value := range authcontroller.AUC {
		routes.HandleFunc(value.Url, value.Control).Methods(value.Method)
	}
	fmt.Println("Import of AUTH routes Completed")

	// USER
	for _, value := range usercontroller.UC {
		routes.HandleFunc(value.Url, value.Control).Methods(value.Method)
	}
	fmt.Println("Import of USER routes Completed")

	// LIST
	for _, value := range listcontroller.LC {
		routes.HandleFunc(value.Url, value.Control).Methods(value.Method)
	}
	fmt.Println("Import of LIST routes Completed")

	// Configuración de CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Cambia esto a tu origen permitido
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	// Aplica CORS como middleware
	handler := corsOptions.Handler(routes)

	fmt.Println("API Listening on port 8080")
	log.Fatal(http.ListenAndServe(envirimoents.GetPort(), handler))
}
