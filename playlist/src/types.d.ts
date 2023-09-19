export type WID<T> = T & {
  id: string;
};

export interface User {
  name: string;
  lastname: string;
  email: string;
  password: string;
}

// type Code struct {
// 	Id           int    `json:"id"`
// 	Code         string `json:"code"`
// 	Order_Number int    `json:"order_number"`
// 	IsPlatey     bool   `json:"isplatey"`
// 	IdUser       string `json:"iduser"`
// 	IdList       string `json:"idlist"`
// }

// var CodeMuckUp Code = Code{}
// var CodeMuckP []Code = []Code{}

// type Login struct {
// 	Email    string
// 	Password string
// }

// type LoginResponse struct {
// 	Data  any    `json:"data"`
// 	Token string `json:"token"`
// }

// type Paginate struct {
// 	Data   any    `json:"data"`
// 	Limit  string `json:"limit"`
// 	Offset string `json:"offset"`
// }

// type Message struct {
// 	Message any `json:"message"`
// }

// // Act === Actual
// type List struct {
// 	Id     int    `json:"id"`
// 	Name   string `json:"name"`
// 	IdUser string `json:"iduser"`
// 	Act    int    `json:"act"`
// 	Counts int    `json:"counts"`
// }

// var ListMuckUp List = List{}
// var ListMuckP []List = []List{}

// type Controller struct {
// 	Url     string
// 	Method  string
// 	Control func(http.ResponseWriter, *http.Request)
// }

// type Iframe struct {
// 	Ifr  string `json:"ifr"`
// 	Type string `json:"type"`
// }

// type User struct {
// 	Id       int16  `json:"id"`
// 	Name     string `json:"name"`
// 	LastName string `json:"lastname"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// var UserMuckUp User = User{}
// var UserMuckP []User = []User{}
