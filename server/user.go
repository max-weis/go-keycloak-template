package server

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	templ "github.com/max-weis/auth-app-template"
	"net/http"
)

type AuthHandler struct {
	Service templ.UserService

	Logger echo.Logger
}

func (a *AuthHandler) Login(c echo.Context) error {
	var user templ.LoginUser
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		a.Logger.Errorf("Could not decode user: %v", err)
		http.Error(c.Response(), "Could not decode user", http.StatusForbidden)
		return err
	}

	a.Logger.Infof("Logging in user: %s", user.Username)

	token, err := a.Service.Login(user)
	if err != nil {
		a.Logger.Errorf("Could not login user: %v", err)
		http.Error(c.Response(), "Could not login user", http.StatusForbidden)
		return err
	}

	if err := json.NewEncoder(c.Response()).Encode(token); err != nil {
		a.Logger.Errorf("Could not encode user: %v", err)
		http.Error(c.Response(), "Could not encode user", http.StatusInternalServerError)
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (a *AuthHandler) Register(c echo.Context) error {
	var user templ.NewUser
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		a.Logger.Errorf("Could not decode user: %v", err)
		http.Error(c.Response(), "Could not decode user", http.StatusInternalServerError)
		return err
	}

	userId, err := a.Service.Register(user)
	if err != nil {
		a.Logger.Errorf("Could not register user: %v", err)
		http.Error(c.Response(), "Could not register user", http.StatusForbidden)
		return err
	}

	return c.String(http.StatusCreated, userId)
}
