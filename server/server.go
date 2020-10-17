package server

import (
	"github.com/Nerzal/gocloak/v7"
	"github.com/labstack/echo/v4"
	templ "github.com/max-weis/auth-app-template"
	"github.com/max-weis/auth-app-template/pkg/config"
)

type Server struct {
	Echo   *echo.Echo
	Client gocloak.GoCloak
}

func NewServer(echo *echo.Echo, client gocloak.GoCloak, service templ.UserService) *Server {
	srv := &Server{Echo: echo, Client: client}

	authHandler := AuthHandler{
		Service: service,
		Logger: echo.Logger,
	}

	srv.Echo.POST("/auth/login", authHandler.Login)
	srv.Echo.POST("/auth/register", authHandler.Register)

	return srv
}

func (s *Server) Serve(config config.Config) {
	s.Echo.Logger.Fatal(s.Echo.Start(config.Port))
}
