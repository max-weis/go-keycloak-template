package main

import (
	"github.com/Nerzal/gocloak/v7"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/max-weis/auth-app-template/pkg/config"
	"github.com/max-weis/auth-app-template/server"
	"github.com/max-weis/auth-app-template/user"
)

func main() {
	cfg := config.NewConfig()

	client := gocloak.NewClient(cfg.KeycloakHost)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userService := &user.Service{
		Logger: e.Logger,
		Client: client,
		Config: cfg,
	}

	srv := server.NewServer(e, client, userService)

	srv.Serve(cfg)
}
