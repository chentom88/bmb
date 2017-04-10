//go:generate hel
package handlers

import (
	"encoding/json"
	"net/http"
	"users"
)

type UserService interface {
	RegisterUser(newUser *users.User) error
	GetUser(emailAddress string) (*users.User, error)
}

type UserHandler struct {
	userService UserService
}

func (u *UserHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	t := &users.User{}
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	u.userService.RegisterUser(t)
}

func NewUserHandler(u UserService) *UserHandler {
	return &UserHandler{
		userService: u,
	}
}
