package types

import (
	"net/http"
)

type Code struct {
	Id       string `json:"id"`
	Code     string `json:"code"`
	Order    string `json:"order"`
	IsPlayed bool   `json:"isplayed"`
	IdUser   string `json:"iduser"`
	IdList   string `json:"idlist"`
}

var CodeMuckUp Code = Code{}
var CodeMuckP []Code = []Code{}

type Paginate struct {
	Data   any    `json:"data"`
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

type Message struct {
	Message any `json:"message"`
}

type List struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	IdUser string `json:"iduser"`
	Act    int    `json:"act"`
}

var ListMuckUp List = List{}
var ListMuckP []List = []List{}

type Controller struct {
	Url     string
	Method  string
	Control func(http.ResponseWriter, *http.Request)
}
