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

type PlayList struct {
	Id     string   `json:"id"`
	Code   string   `json:"code"`
	IdUser string   `json:"iduser"`
	IdList []string `json:"idlist"`
}

type Controller struct {
	Url     string
	Method  string
	Control func(http.ResponseWriter, *http.Request)
}
