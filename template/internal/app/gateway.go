package app

import (
	context "context"

	_ "{{ if Owner }}{{ Owner }}{{ end }}.{{ Project }}/api/proto/{{ Project }}/api"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	grpc "google.golang.org/grpc"
)

// Gateway ...
type Gateway = gwruntime.ServeMux

// NewGateway ...
func NewGateway(ctx context.Context) *Gateway {
	return gwruntime.NewServeMux()
}

// InitGateway ...
func (app *App) InitGateway(ctx context.Context) error {
	options := []grpc.DialOption{grpc.WithInsecure()}

	for endpoint, fn := range map[string]func(context.Context, *gwruntime.ServeMux, string, []grpc.DialOption) error{
		// app.config.system.account.address: proto.RegisterAccountServiceHandlerFromEndpoint,
		// app.config.system.auth.address:    proto.RegisterAuthServiceHandlerFromEndpoint,
		// app.config.system.media.address:   proto.RegisterMediaServiceHandlerFromEndpoint,
	} {
		if err := fn(ctx, app.gateway, endpoint, options); err != nil {
			return err
		}
	}

	return nil
}
