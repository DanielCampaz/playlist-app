package listcontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	list "main/src/services/List"
	"main/src/types"
	"main/src/utils"
	"net/http"

	"github.com/gorilla/mux"
)

var endpoint = utils.CreateEndpointControllers("list")

var get = endpoint("{id}")
var methodGet = "GET"

func getList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listId := vars["id"]

	list, _ := list.GetList(listId)

	utils.JsonResponse(w, list)
}

var getAll = endpoint("g/all")
var methodGetAll = "GET"

func getAllLists(w http.ResponseWriter, r *http.Request) {
	// Obtener los parámetros de consulta de la URL
	queryParams := r.URL.Query()

	// Obtener valores individuales de los parámetros de consulta
	limit := queryParams.Get("limit")
	offset := queryParams.Get("offset")

	lists, _ := list.GetLists(limit, offset)

	utils.JsonResponse(w, lists)
}

var create = endpoint("create")
var methodCreate = "POST"

func createList(w http.ResponseWriter, r *http.Request) {

	// Leer el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Deserializar el cuerpo en una estructura User
	var newList types.List
	err = json.Unmarshal(body, &newList)
	if err != nil {
		http.Error(w, "Error al deserializar los datos del cuerpo", http.StatusBadRequest)
		return
	}

	ersr := list.CreateList(newList)
	if ersr != nil {
		utils.JsonResponse(w, ersr)

	} else {
		utils.JsonResponse(w, newList)
	}

}

var update = endpoint("update/{id}")
var methodUpdate = "PUT"

func updateList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID := vars["id"]
	lista, err := list.GetList(listID)
	if err != nil {
		// TODO: Enviar al correo
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Deserializar el cuerpo en una estructura User
	var upList types.List = lista
	err = json.Unmarshal(body, &upList)
	if err != nil {
		http.Error(w, "Error al deserializar los datos del cuerpo", http.StatusBadRequest)
		return
	}

	errs := list.UpdateList(listID, upList)
	if errs != nil {
		// TODO: Enviar al correo
		fmt.Print(errs)
	}

	utils.JsonResponse(w, upList)
}

var delete = endpoint("delete/{id}")
var methodDelete = "DELETE"

func deleteList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID := vars["id"]
	err := list.DeleteList(listID)
	if err != nil {
		// TODO: Enviar al correo
	}

	// TODO: Eliminar los codigos
	utils.JsonResponse(w, types.Message{Message: "List Deleting"})
}

var LC []types.Controller = []types.Controller{
	{
		Url:     create,
		Control: createList,
		Method:  methodCreate,
	},
	{
		Url:     update,
		Control: updateList,
		Method:  methodUpdate,
	},
	{
		Url:     delete,
		Control: deleteList,
		Method:  methodDelete,
	},
	{
		Url:     get,
		Control: getList,
		Method:  methodGet,
	},
	{
		Url:     getAll,
		Control: getAllLists,
		Method:  methodGetAll,
	},
}
