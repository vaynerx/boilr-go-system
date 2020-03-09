package app

import (
	context "context"
)

// App ...
type App struct {
	config  *Config
	gateway *Gateway
	router  *Router
}

// NewApp ...
func NewApp(ctx context.Context) *App {
	return &App{
		skaffold: NewSkaffold(ctx),
		config:   NewConfig(ctx),
		router:   NewRouter(ctx),
	}
}

// InitApp ...
func (app *App) InitApp(ctx context.Context) error {
	var err error

	if err = app.InitConfig(ctx); err != nil {
		return err
	}

	if err = app.InitGRPC(ctx); err != nil {
		return err
	}

	return nil
}
