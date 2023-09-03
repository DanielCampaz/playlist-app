package utils

import (
	"encoding/json"
	"fmt"
	ENV "main/src/envirimoents"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func GetUrlMysqlConnection() string {

	user := ENV.GetEnv("MYSQL_USER")
	password := ENV.GetEnv("MYSQL_PASSWORD")
	port := ENV.GetEnv("MYSQL_PORT")
	namedatabase := ENV.GetEnv("MYSQL_NAMEDATABASE")

	if user == "NULL" || password == "NULL" || port == "NULL" || namedatabase == "NULL" {
		return ""
	}

	url := user + ":" + password + "@tcp(localhost:" + port + ")/" + namedatabase

	return url
}

func GetApiName() string {

	apiName := ENV.GetEnv("API_NAME")
	if apiName == "NULL" {
		return ""
	}
	return apiName + "/"
}

func CreateEndpoint(endpointController string, endpoint string) string {
	endpointe := "/" + GetApiName() + endpointController + "/" + endpoint
	fmt.Println(endpointe)
	return endpointe
}

func CreateEndpointControllers(endpointController string) func(endpoint string) string {
	return func(endpoint string) string {
		return CreateEndpoint(endpointController, endpoint)
	}
}

func JsonResponse(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(v)
}

func HashPassword(password string) (string, error) {
	// Genera un hash bcrypt a partir de la contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(inputPassword string, hashedPassword string) error {
	// Compara la contraseña ingresada con el hash almacenado
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}

var Tables []string = []string{
	"CREATE TABLE IF NOT EXISTS users (id INT NOT NULL AUTO_INCREMENT,name VARCHAR(200) NOT NULL,email VARCHAR(200) NOT NULL,password VARCHAR(200) NOT NULL,lastname VARCHAR(200) NOT NULL,PRIMARY KEY (id));",
	"CREATE TABLE IF NOT EXISTS code (id INT NOT NULL AUTO_INCREMENT,code VARCHAR(200) NOT NULL,order INT NOT NULL,isplatey TINYINT NOT NULL,iduser INT NOT NULL,idlist INT NOT NULL,PRIMARY KEY (id));",
	"CREATE TABLE IF NOT EXISTS list (id INT NOT NULL AUTO_INCREMENT,name VARCHAR(200) NOT NULL,iduser INT NOT NULL,act INT NULL DEFAULT 0,PRIMARY KEY (id));",
}
