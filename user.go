package templ

import "github.com/Nerzal/gocloak/v7"

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserService interface {
	Login(user LoginUser) (*gocloak.JWT, error)
	Register(user NewUser) (string, error)
}
