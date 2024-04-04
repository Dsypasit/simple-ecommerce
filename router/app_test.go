package router

import (
	"testing"

	"github.com/Dsypasit/simple-ecommerce/config"
	"github.com/stretchr/testify/assert"
)

type MockRegister struct{}

func (r *MockRegister) Register(app *App) error { return nil }

func TestRouteRegister(t *testing.T) {
	test := struct {
		name      string
		mockRoute Routes
	}{
		name:      "get no error when register route",
		mockRoute: &MockRegister{},
	}

	t.Run(test.name, func(t *testing.T) {
		cfg := config.New()
		app := New(cfg)

		err := app.RegisterRoute(test.mockRoute)
		assert.NoError(t, err)
	})
}
