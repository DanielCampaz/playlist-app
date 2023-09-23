package utils

import (
	"encoding/json"
	"fmt"
	ENV "main/src/envirimoents"
	"main/src/types"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func GetUrlMysqlConnection() string {

	user := ENV.GetEnv("MYSQL_USER", "root")
	password := ENV.GetEnv("MYSQL_PASSWORD", "123456789")
	port := ENV.GetEnv("MYSQL_PORT", "3306")
	namedatabase := ENV.GetEnv("MYSQL_NAMEDATABASE", "playlist")

	if user == "NULL" || password == "NULL" || port == "NULL" || namedatabase == "NULL" {
		return ""
	}

	url := user + ":" + password + "@tcp(localhost:" + port + ")/" + namedatabase

	return url
}

func GetSecretKeyJWT() string {
	nameEnv := "SECRETKEYJWT"
	sk := ENV.GetEnv(nameEnv, nameEnv)
	return sk
}

func GetApiName() string {

	apiName := ENV.GetEnv("API_NAME", "")
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

func VerifyIfIsEncrypt(password string) bool {
	long := len(password)
	fmt.Print(long)
	if long > 30 {
		return true
	} else {
		return false
	}
}

var Tables []string = []string{
	"CREATE TABLE IF NOT EXISTS users (id INT NOT NULL AUTO_INCREMENT,name VARCHAR(200) NOT NULL,email VARCHAR(200) NOT NULL,password VARCHAR(200) NOT NULL,lastname VARCHAR(200) NOT NULL,PRIMARY KEY (id));",
	"CREATE TABLE IF NOT EXISTS code (id INT NOT NULL AUTO_INCREMENT,code VARCHAR(200) NOT NULL,order_number INT NOT NULL,isplatey TINYINT NOT NULL,iduser INT NOT NULL,idlist INT NOT NULL,PRIMARY KEY (id));",
	"CREATE TABLE IF NOT EXISTS list (id INT NOT NULL AUTO_INCREMENT,name VARCHAR(200) NOT NULL,iduser INT NOT NULL,act INT NULL DEFAULT 0,PRIMARY KEY (id));",
}

func IframeRemove(iframeCode types.Iframe) string {
	// Define una expresión regular para buscar el código del video de YouTube
	regex := regexp.MustCompile(`https://www\.youtube\.com/embed/([A-Za-z0-9_-]+)`)

	if iframeCode.Type == "iframe" {
		// Busca coincidencias en la cadena del iframe
		matches := regex.FindStringSubmatch(iframeCode.Ifr)

		// Comprueba si se encontraron coincidencias
		if len(matches) >= 2 {
			videoCode := matches[1]
			return videoCode
		} else {
			return ""
		}

	} else if iframeCode.Type == "url" {
		re := regexp.MustCompile(`[?&]v=([a-zA-Z0-9_-]+)|youtu\.be/([a-zA-Z0-9_-]+)`)

		// Encuentra todas las coincidencias en la URL
		matches := re.FindStringSubmatch(iframeCode.Ifr)

		// Verifica si se encontraron coincidencias
		if len(matches) < 2 {
			fmt.Print("Could not find video code in URL")
			return ""
		}

		// El código de video puede estar en la primera o segunda subcoincidencia
		for i := 1; i < len(matches); i++ {
			if matches[i] != "" {
				return matches[i]
			}
		}
		fmt.Print("Could not find video code in URL")

		return ""
	} else {
		return iframeCode.Ifr
	}

}
