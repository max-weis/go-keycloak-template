package user

import (
	"context"
	"github.com/Nerzal/gocloak/v7"
	"github.com/labstack/echo/v4"
	"github.com/max-weis/auth-app-template"
	"github.com/max-weis/auth-app-template/pkg/config"
)

type Service struct {
	Logger echo.Logger
	Client gocloak.GoCloak
	Config config.Config
}

func (s *Service) Login(user templ.LoginUser) (*gocloak.JWT, error) {
	token, err := s.Client.Login(
		context.TODO(),
		s.Config.ClientId,
		s.Config.ClientSecret,
		s.Config.Realm,
		user.Username,
		user.Password,
	)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *Service) Register(user templ.NewUser) (string, error) {
	token, err := s.Client.LoginAdmin(
		context.TODO(),
		s.Config.AdminUser,
		s.Config.AdminPass,
		s.Config.Realm,
	)

	if err != nil {
		return "", err
	}

	enabled := true
	newUser := gocloak.User{
		Email:         &user.Email,
		Username:      &user.Username,
		Enabled:       &enabled,
		EmailVerified: &enabled,
	}

	userId, err := s.Client.CreateUser(context.TODO(), token.AccessToken, s.Config.Realm, newUser)
	if err != nil {
		return "", err
	}

	err = s.Client.SetPassword(context.TODO(), token.AccessToken, userId, s.Config.Realm, user.Password, false)
	if err != nil {
		return "", err
	}

	return userId, nil
}
