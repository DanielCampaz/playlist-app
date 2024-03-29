package usercontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	u "main/src/services/User"
	"main/src/types"
	"main/src/utils"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

var endpoint = utils.CreateEndpointControllers("users")

var get = endpoint("{id}")
var methodGet = "GET"

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := u.GetUser(userID)

	if err != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error to Get user"})
		return
	} else {

		utils.JsonResponse(w, user)
	}

}

var getAll = endpoint("g/all")
var methodGetAll = "GET"

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	// Obtener los parámetros de consulta de la URL
	queryParams := r.URL.Query()

	// Obtener valores individuales de los parámetros de consulta
	limit := queryParams.Get("limit")
	offset := queryParams.Get("offset")

	users, err := u.GetUsers(limit, offset)

	if err != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error to Get all user"})
		return
	} else {

		utils.JsonResponse(w, users)
	}

}

var getId = endpoint("g/{id}")
var methodGetId = "GET"

func getIdUsers(w http.ResponseWriter, r *http.Request) {
	// Obtener los parámetros de consulta de la URL
	queryParams := r.URL.Query()

	// Obtener valores individuales de los parámetros de consulta
	limit := queryParams.Get("limit")
	offset := queryParams.Get("offset")

	users, err := u.GetUsers(limit, offset)

	if err != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error to Get user"})
		return
	} else {

		utils.JsonResponse(w, users)
	}

}

var create = endpoint("create")
var methodCreate = "POST"

func CreateUser(w http.ResponseWriter, r *http.Request) {

	// Leer el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Deserializar el cuerpo en una estructura User
	var newUser types.User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, "Error deserializing body data", http.StatusBadRequest)
		return
	}

	ersr := u.CreateUser(newUser)
	if ersr != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error to create user"})

	} else {
		newUserCreate, err := u.GetUserByEmail(newUser.Email)
		if err != nil {
			utils.JsonResponse(w, types.ErrorMessage{Error: "Error to Get user"})
			return

		} else {
			utils.JsonResponse(w, newUserCreate)
		}
	}

}

var update = endpoint("update/{id}")
var methodUpdate = "PUT"

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user, err := u.GetUser(userID)
	if err != nil {
		// TODO: Enviar al correo
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Deserializar el cuerpo en una estructura User
	var upUser types.User = user
	err = json.Unmarshal(body, &upUser)
	if err != nil {
		http.Error(w, "Error deserializing body data", http.StatusBadRequest)
		return
	}

	errs := u.UpdateUser(userID, upUser)
	if errs != nil {
		// TODO: Enviar al correo
		fmt.Print(errs)
		return
	}

	id := strconv.FormatInt(int64(upUser.Id), 10)
	userUpdate, err := u.GetUser(id)
	if err != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error to Update user"})
		return

	} else {
		utils.JsonResponse(w, userUpdate)
	}
}

var delete = endpoint("delete/{id}")
var methodDelete = "DELETE"

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	err := u.DeleteUser(userID)
	if err != nil {
		// TODO: Enviar al correo
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error to Delete user"})
		return

	}
	utils.JsonResponse(w, types.Message{Message: "User Deleting"})
}

var UC []types.Controller = []types.Controller{
	{
		Url:     create,
		Control: utils.VerifyTokenJWT(CreateUser),
		Method:  methodCreate,
	},
	{
		Url:     update,
		Control: utils.VerifyTokenJWT(updateUser),
		Method:  methodUpdate,
	},
	{
		Url:     delete,
		Control: utils.VerifyTokenJWT(deleteUser),
		Method:  methodDelete,
	},
	{
		Url:     get,
		Control: utils.VerifyTokenJWT(getUser),
		Method:  methodGet,
	},
	{
		Url:     getAll,
		Control: utils.VerifyTokenJWT(getAllUsers),
		Method:  methodGetAll,
	},
	{
		Url:     getId,
		Control: utils.VerifyTokenJWT(getIdUsers),
		Method:  methodGetId,
	},
}
