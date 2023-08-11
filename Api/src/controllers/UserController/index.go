package usercontroller

import (
	"main/src/types"
	"main/src/utils"
	"net/http"
)

var endpoint = utils.CreateEndpointControllers("users")

var create = endpoint("create")

func createUser(w http.ResponseWriter, r *http.Request) {
	user := types.User{Name: "***", LastName: "***", Email: "***@***.com", Password: "***"}
	utils.JsonResponse(w, user)
}

var UC []types.Controller = []types.Controller{
	{
		Url:     create,
		Control: createUser,
	},
}
