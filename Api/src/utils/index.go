package utils

import (
	"encoding/json"
	"fmt"
	ENV "main/src/envirimoents"
	"net/http"
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
