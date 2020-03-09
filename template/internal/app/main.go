package app

import (
	context "context"
)

// App ...
type App struct {
	skaffold *Skaffold
	config   *Config
	gateway  *Gateway
	router   *Router
}

// NewApp ...
func NewApp(ctx context.Context) *App {
	return &App{
		skaffold: NewSkaffold(ctx),
		config:   NewConfig(ctx),
		gateway:  NewGateway(ctx),
		router:   NewRouter(ctx),
	}
}

// InitApp ...
func (app *App) InitApp(ctx context.Context) error {
	var err error

	if err = app.InitSkaffold(ctx); err != nil {
		return err
	}

	if err = app.InitConfig(ctx); err != nil {
		return err
	}

	if err = app.InitGateway(ctx); err != nil {
		return err
	}

	if err = app.InitRouter(ctx); err != nil {
		return err
	}

	return nil
}
