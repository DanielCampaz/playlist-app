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

	list, err := list.GetList(listId)

	if err != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error to Get List"})
		return
	} else {

		utils.JsonResponse(w, list)
	}

}

var getAll = endpoint("g/all")
var methodGetAll = "GET"

func getAllLists(w http.ResponseWriter, r *http.Request) {
	// Obtener los parámetros de consulta de la URL
	queryParams := r.URL.Query()

	// Obtener valores individuales de los parámetros de consulta
	limit := queryParams.Get("limit")
	offset := queryParams.Get("offset")

	lists, err := list.GetLists(limit, offset)

	if err != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error to Get All List"})
		return
	} else {

		utils.JsonResponse(w, lists)
	}

}

var getCodes = endpoint("codes/all/{idlist}")
var methodGetCodes = "GET"

func getAllCodes(w http.ResponseWriter, r *http.Request) {
	// Obtener los parámetros de consulta de la URL
	queryParams := r.URL.Query()
	vars := mux.Vars(r)
	listID := vars["idlist"]

	// Obtener valores individuales de los parámetros de consulta
	limit := queryParams.Get("limit")
	offset := queryParams.Get("offset")

	lists, err := code.GetCodesByList(limit, offset, listID)

	if err != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error to Get All Code to " + listID})
		return
	} else {

		utils.JsonResponse(w, lists)
	}

}

var create = endpoint("create")
var methodCreate = "POST"

func createList(w http.ResponseWriter, r *http.Request) {

	// Leer el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Deserializar el cuerpo en una estructura User
	var newList types.List
	err = json.Unmarshal(body, &newList)
	if err != nil {
		fmt.Println("Error deserializing JSON:", err)
		http.Error(w, "Error deserializing request body", http.StatusBadRequest)
		return
	}

	newList.Act = 0
	newList.Counts = 0
	ersr := list.CreateList(newList)
	if ersr != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error To Create List"})
		return
	} else {
		listG, err := list.GetListByName(newList.Name)
		if err != nil {
			utils.JsonResponse(w, types.ErrorMessage{Error: "Error to Get List"})
			return

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
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Deserializar el cuerpo en una estructura User
	var upList types.List = lista
	err = json.Unmarshal(body, &upList)
	if err != nil {
		http.Error(w, "Error deserializing body data", http.StatusBadRequest)
		return
	}

	errs := list.UpdateList(listID, upList)
	if errs != nil {
		// TODO: Enviar al correo
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error to Update List"})
		return

	}

	listar, errv := list.GetList(listID)
	if errv != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error to Get List"})
		return

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
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error to Delete List"})
		return
	}

	codeDel := code.DeleteAllCodes(listID)
	msDelCod := "Error when removing code from this playlist: " + listID
	if codeDel {
		msDelCod = "All Code Removed"
	}
	// TODO: Eliminar los codigos
	utils.JsonResponse(w, types.Message{Message: "List Deleting; " + msDelCod})
}

var next = endpoint("next/{idList}/{idUser}")
var methodNext = "GET"

func nextCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID := vars["idList"]
	userID := vars["idUser"]
	lista, err := list.GetList(listID)
	if err != nil {
		utils.JsonResponse(w, types.Message{Message: "Error getting the List"})
		return
	}

	_, errc := user.GetUser(userID)
	if errc != nil {
		utils.JsonResponse(w, "Error getting user")
		return
	}

	act := lista.Act + 1
	if act > lista.Counts {
		utils.JsonResponse(w, types.Message{Message: "There are no more videos in the list"})
		return
	}
	codes, errors := code.GetCodesByOrder(listID, fmt.Sprint(act))
	if errors != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error getting current code"})
		return
	}

	errorror := "Error updating list"
	codes[0].IsPlatey = true
	erroo := code.UpdateCode(codes[0].Id, codes[0])
	if erroo != nil {
		errorror = errorror + "; Error updating code"
	}

	ee := list.UpdateAct(listID)
	if ee != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: errorror})
		return
	} else {
		utils.JsonResponse(w, codes)
	}

}

var add = endpoint("add/{idList}/{idUser}")
var methodAdd = "POST"

func addCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID := vars["idList"]
	userID := vars["idUser"]
	lista, err := list.GetList(listID)
	if err != nil {
		utils.JsonResponse(w, "Error getting the List")
		return
	}

	_, errc := user.GetUser(userID)
	if errc != nil {
		utils.JsonResponse(w, "Error getting the User")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Deserializar el cuerpo en una estructura User

	var iframe types.Iframe
	err = json.Unmarshal(body, &iframe)
	if err != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error when opting for Body"})
		return
	}

	fr := utils.IframeRemove(iframe)

	if fr == "" {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error when copying the video code"})
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

	errs := code.CreateCode(newCode)
	if errs != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error creating Code"})
		return
	}
	erre := list.UpdateCount(listID)

	if erre != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error to update count"})
		return
	} else {
		utils.JsonResponse(w, types.Message{Message: "Code Create Succesful"})
	}

}

var deleteMany = endpoint("deletemany")
var methodDeleteMany = "POST"

func deleteManyCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["idUser"]

	_, errc := user.GetUser(userID)
	if errc != nil {
		utils.JsonResponse(w, "Error getting user")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Deserializar el cuerpo en una estructura User

	var deletemany types.DeleteMany
	err = json.Unmarshal(body, &deletemany)
	if err != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error reading body"})
		return
	}

	dle := list.DeleteManyList(deletemany.Ids)

	utils.JsonResponse(w, dle)
	return

}

var restarte = endpoint("restar/{id}")
var methodRestarte = "PUT"

func restart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	dle := list.UpdateReset(id)
	if dle != nil {
		utils.JsonResponse(w, types.ErrorMessage{Error: "Error to Reset Playlist"})
		return
	} else {
		utils.JsonResponse(w, types.Message{Message: "Reset Complete"})
		return
	}

}

var LC []types.Controller = []types.Controller{
	{
		Url:     create,
		Control: utils.VerifyTokenJWT(createList),
		Method:  methodCreate,
	},
	{
		Url:     restarte,
		Control: utils.VerifyTokenJWT(restart),
		Method:  methodRestarte,
	},
	{
		Url:     update,
		Control: utils.VerifyTokenJWT(updateList),
		Method:  methodUpdate,
	},
	{
		Url:     delete,
		Control: utils.VerifyTokenJWT(deleteList),
		Method:  methodDelete,
	},
	{
		Url:     deleteMany,
		Control: utils.VerifyTokenJWT(deleteManyCode),
		Method:  methodDeleteMany,
	},
	{
		Url:     get,
		Control: utils.VerifyTokenJWT(getList),
		Method:  methodGet,
	},
	{
		Url:     getAll,
		Control: utils.VerifyTokenJWT(getAllLists),
		Method:  methodGetAll,
	},
	{
		Url:     next,
		Control: utils.VerifyTokenJWT(nextCode),
		Method:  methodNext,
	},
	{
		Url:     add,
		Control: utils.VerifyTokenJWT(addCode),
		Method:  methodAdd,
	},
	{
		Url:     getCodes,
		Control: utils.VerifyTokenJWT(getAllCodes),
		Method:  methodGetCodes,
	},
}
