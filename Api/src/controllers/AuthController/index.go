package authcontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	usercontroller "main/src/controllers/UserController"

	//u "main/src/services/User"
	"main/src/types"
	"main/src/utils"
	"net/http"
)

var endpoint = utils.CreateEndpointControllers("auth")

var login = endpoint("login")
var methodLogin = "POST"

func loginM(w http.ResponseWriter, r *http.Request) {
	// Leer el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Deserializar el cuerpo en una estructura User
	var login types.Login
	err = json.Unmarshal(body, &login)
	if err != nil {
		http.Error(w, "Error al deserializar los datos del cuerpo", http.StatusBadRequest)
		return
	}

	fmt.Print(login)

	// user, err := u.GetUserByEmail(login.Email)
	// if err != nil {
	// 	utils.JsonResponse(w, err)
	// } else {
	// 	eror := utils.ComparePassword(login.Password, user.Password)
	// 	if eror != nil {
	// 		utils.JsonResponse(w, types.Message{Message: "password is not equal"})
	// 	} else {
	// 		token, errorr := utils.GenerateJWT(user.Email)
	// 		if errorr != nil {
	// 			utils.JsonResponse(w, types.LoginResponse{Data: "Error to generate Token", Token: ""})
	// 		}
	// 		utils.JsonResponse(w, types.LoginResponse{Data: user, Token: token})
	// 	}
	// }
}

var singup = endpoint("singup")
var methodSingup = "POST"

func singupM(w http.ResponseWriter, r *http.Request) {

	usercontroller.CreateUser(w, r)
}

// var generar = endpoint("generate")
// var methodGenerar = "GET"

// func generarEnd(w http.ResponseWriter, r *http.Request) {
// 	token, err := utils.GenerateJWT("jcbakjsbc@gmail.com")
// 	if err != nil {
// 		utils.JsonResponse(w, types.Message{Message: err})
// 	} else {
// 		utils.JsonResponse(w, types.Message{Message: token})
// 	}
// }

// var verificar = endpoint("verificar")
// var methodVerificar = "GET"

// func verificarEnd(w http.ResponseWriter, r *http.Request) {
// 	utils.JsonResponse(w, types.Message{Message: "Token Verificado"})
// }

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

/*
{
		Url:     generar,
		Method:  methodGenerar,
		Control: generarEnd,
	}, {
		Url:     verificar,
		Method:  methodVerificar,
		Control: utils.VerifyTokenJWT(verificarEnd),
	},


*/
