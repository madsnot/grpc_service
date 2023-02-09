package server

import (
	"context"
	"log"
	"net"

	"github.com/madsnot/grpc_service/grpc/api"
	"github.com/madsnot/grpc_service/pkg/errors"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

const servDirPath = "C:/Images"

type GRPCServer struct {
	api.ImagesHandlerServer
}

func Run() {
	unaryLimiter := rate.NewLimiter(100, 100)
	streamLimiter := rate.NewLimiter(10, 10)

	serv := grpc.NewServer(
		grpc.StreamInterceptor(func(srv interface{}, stream grpc.ServerStream,
			info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

			if streamLimiter.Allow() {
				return handler(srv, stream)
			}

			return errors.ConnLimitExceededError{Method: info.FullMethod}.Error()
		}),

		grpc.UnaryInterceptor(func(ctx context.Context, req interface{},
			info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

			if unaryLimiter.Allow() {
				return handler(ctx, req)
			}

			return nil, errors.ConnLimitExceededError{Method: info.FullMethod}.Error()
		}),
	)

	grpcServer := &GRPCServer{}
	api.RegisterImagesHandlerServer(serv, grpcServer)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
		return
	}

	if err := serv.Serve(listener); err != nil {
		log.Println(err)
		return
	}
}
