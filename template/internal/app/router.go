package app

import (
	context "context"

	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

// Router ...
type Router = echo.Echo

// NewRouter ...
func NewRouter(ctx context.Context) *Router {
	return echo.New()
}

// InitRouter ...
func (app *App) InitRouter(ctx context.Context) error {
	app.router.Use(middleware.Secure())
	app.router.Use(middleware.Logger())
	app.router.Use(middleware.Recover())
	app.router.Use(middleware.CORS())
	app.router.GET("/", echo.WrapHandler(app.gateway))

	return app.router.Start(":" + app.config.api.{{ Project }}.Port)
}
