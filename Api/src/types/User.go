package types

type User struct {
	Id       int16  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var UserMuckUp User = User{}
var UserMuckP []User = []User{}
