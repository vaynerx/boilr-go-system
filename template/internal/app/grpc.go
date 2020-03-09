package app

import (
	context "context"
	fmt "fmt"
	net "net"

	proto "{{ Owner }}.{{ Project }}/api/proto/{{ Project }}/api"

	ocgrpc "go.opencensus.io/plugin/ocgrpc"
	grpc "google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	reflection "google.golang.org/grpc/reflection"

	log "github.com/sirupsen/logrus"
)

// InitGPRC ...
func (app *App) InitGPRC(ctx context.Context) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", local.server.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))

	proto.RegisterAccountServiceServer(grpcServer, app)
	reflection.Register(grpcServer)
	healthpb.RegisterHealthServer(grpcServer, app)

	log.Infof("starting to listen on tcp: %q", listen.Addr().String())
	return grpcServer.Serve(listen)
}
