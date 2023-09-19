package types

import (
	"net/http"
)

type Code struct {
	Id           int    `json:"id"`
	Code         string `json:"code"`
	Order_Number int    `json:"order_number"`
	IsPlatey     bool   `json:"isplatey"`
	IdUser       string `json:"iduser"`
	IdList       string `json:"idlist"`
}

var CodeMuckUp Code = Code{}
var CodeMuckP []Code = []Code{}

type Login struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Data  any    `json:"data"`
	Token string `json:"token"`
}

type Paginate struct {
	Data   any    `json:"data"`
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

type Message struct {
	Message any `json:"message"`
}

type ErrorMessage struct {
	Error any `json:"error"`
}

// Act === Actual
type List struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	IdUser string `json:"iduser"`
	Act    int    `json:"act"`
	Counts int    `json:"counts"`
}

var ListMuckUp List = List{}
var ListMuckP []List = []List{}

type Controller struct {
	Url     string
	Method  string
	Control func(http.ResponseWriter, *http.Request)
}

type Iframe struct {
	Ifr  string `json:"ifr"`
	Type string `json:"type"`
}
