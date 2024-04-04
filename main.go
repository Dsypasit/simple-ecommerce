package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Dsypasit/simple-ecommerce/config"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg := config.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(fmt.Sprintf(":%v", cfg.Server.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down server")
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Server.ShutdownTimeout)*time.Second)

	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
