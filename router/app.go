package router

import (
	"github.com/Dsypasit/simple-ecommerce/config"
	"github.com/labstack/echo/v4"
)

type App struct {
	echo *echo.Echo
	cfg  config.Config
}

type Routes interface {
	Register(r *App) error
}

func New(c config.Config) *App {
	return &App{echo: echo.New(), cfg: c}
}

func (a *App) RegisterRoute(routes ...Routes) error {
	for _, route := range routes {
		err := route.Register(a)
		if err != nil {
			return err
		}
	}
	return nil
}
