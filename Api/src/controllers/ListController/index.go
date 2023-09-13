package listcontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	code "main/src/services/Code"
	list "main/src/services/List"
	user "main/src/services/User"
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
		fmt.Println("Error al deserializar JSON:", err)
		http.Error(w, "Error al deserializar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	ersr := list.CreateList(newList)
	if ersr != nil {
		utils.JsonResponse(w, "Error To Create List")
	} else {
		listG, err := list.GetListByName(newList.Name)
		if err != nil {
			utils.JsonResponse(w, err)

		} else {
			utils.JsonResponse(w, listG)
		}
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

	listar, errv := list.GetList(listID)
	if errv != nil {
		utils.JsonResponse(w, errv)
	} else {
		utils.JsonResponse(w, listar)
	}
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

	codeDel := code.DeleteAllCodes(listID)
	msDelCod := "Error al eliminar code's de esta playlist: " + listID
	if codeDel {
		msDelCod = "Todos los Code Eliminados"
	}
	// TODO: Eliminar los codigos
	utils.JsonResponse(w, types.Message{Message: "List Deleting; " + msDelCod})
}

var next = endpoint("next/{id}")
var methodNext = "GET"

func nextCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID := vars["id"]
	lista, err := list.GetList(listID)
	if err != nil {
		utils.JsonResponse(w, types.Message{Message: "Error al optener La Lista"})
		return
	}

	act := lista.Act + 1

	codes, errors := code.GetCodesByOrder(listID, fmt.Sprint(act))
	if errors != nil {
		utils.JsonResponse(w, types.Message{Message: "Error al optener Codigo actual"})
		return
	}

	list.UpdateAct(listID)

	utils.JsonResponse(w, codes)
}

var Add = endpoint("add/{idList}/{idUser}")
var methodAdd = "POST"

func addCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID := vars["idList"]
	userID := vars["idUser"]
	lista, err := list.GetList(listID)
	if err != nil {
		utils.JsonResponse(w, "Error al optener la lista")
		return
	}

	_, errc := user.GetUser(userID)
	if errc != nil {
		utils.JsonResponse(w, "Error al optener el usuario")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Deserializar el cuerpo en una estructura User

	var iframe types.Iframe
	err = json.Unmarshal(body, &iframe)
	if err != nil {
		utils.JsonResponse(w, types.Message{Message: err})
		return
	}

	fr := utils.IframeRemove(iframe)

	if fr == "" {
		utils.JsonResponse(w, types.Message{Message: "Error al copiar el codigo del video"})
		return
	}
	// Deserializar el cuerpo en una estructura User
	var newCode types.Code = types.Code{
		IdUser:       userID,
		IdList:       listID,
		Order_Number: lista.Counts + 1,
		IsPlatey:     false,
		Code:         fr,
		Id:           0,
	}

	if err != nil {
		http.Error(w, "Error al deserializar los datos del cuerpo", http.StatusBadRequest)
		return
	}

	errs := code.CreateCode(newCode)
	if errs != nil {
		http.Error(w, "Error al crear Code", http.StatusBadRequest)
		return
	}
	list.UpdateCount(listID)

	utils.JsonResponse(w, types.Message{Message: "Code Creado correctamente"})
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
	{
		Url:     next,
		Control: nextCode,
		Method:  methodNext,
	},
	{
		Url:     Add,
		Control: addCode,
		Method:  methodAdd,
	},
}
