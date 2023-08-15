package authcontroller

import (
	"main/src/types"
	"main/src/utils"
	"net/http"
)

var endpoint = utils.CreateEndpointControllers("auth")

var login = endpoint("login")
var methodLogin = "POST"

func loginM(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, types.Message{Message: "Login"})
}

var singup = endpoint("singup")
var methodSingup = "POST"

func singupM(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, types.Message{Message: "SingUp"})
}

var AUC []types.Controller = []types.Controller{
	{
		Url:     login,
		Method:  methodLogin,
		Control: loginM,
	}, {
		Url:     singup,
		Method:  methodSingup,
		Control: singupM,
	},
}
