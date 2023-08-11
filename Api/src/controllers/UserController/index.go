package usercontroller

import (
	"encoding/json"
	"io/ioutil"
	"main/src/types"
	"main/src/utils"
	"net/http"
)

var endpoint = utils.CreateEndpointControllers("users")

var get = endpoint("")
var methodGet = "GET"

func getUser(w http.ResponseWriter, r *http.Request) {
	user := types.User{Name: "***", LastName: "***", Email: "***@***.com", Password: "***"}
	utils.JsonResponse(w, user)
}

var create = endpoint("create")
var methodCreate = "POST"

func createUser(w http.ResponseWriter, r *http.Request) {

	// Leer el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Deserializar el cuerpo en una estructura User
	var newUser types.User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, "Error al deserializar los datos del cuerpo", http.StatusBadRequest)
		return
	}

	ersr := newUser.CreateUser()
	if ersr != nil {
		utils.JsonResponse(w, ersr)

	} else {
		utils.JsonResponse(w, newUser)
	}

}

var update = endpoint("update")
var methodUpdate = "PUT"

func updateUser(w http.ResponseWriter, r *http.Request) {
	user := types.User{Name: "***", LastName: "***", Email: "***@***.com", Password: "***"}
	utils.JsonResponse(w, user)
}

var delete = endpoint("delete")
var methodDelete = "DELETE"

func deleteUser(w http.ResponseWriter, r *http.Request) {
	user := types.User{Name: "***", LastName: "***", Email: "***@***.com", Password: "***"}
	utils.JsonResponse(w, user)
}

var UC []types.Controller = []types.Controller{
	{
		Url:     create,
		Control: createUser,
		Method:  methodCreate,
	},
	{
		Url:     update,
		Control: updateUser,
		Method:  methodUpdate,
	},
	{
		Url:     delete,
		Control: deleteUser,
		Method:  methodDelete,
	},
	{
		Url:     get,
		Control: getUser,
		Method:  methodGet,
	},
}
