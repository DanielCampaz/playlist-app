package types

import (
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Code struct {
	Code     string `json:"code"`
	Order    string `json:"order"`
	IsPlayed bool   `json:"isplayed"`
	IdUser   string `json:"iduser"`
	IdList   string `json:"idlist"`
}

type PlayList struct {
	Code   string   `json:"code"`
	IdUser string   `json:"iduser"`
	IdList []string `json:"idlist"`
}

type Controller struct {
	Url     string
	Control func(http.ResponseWriter, *http.Request)
}
