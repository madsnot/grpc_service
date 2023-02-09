package server

import (
	"context"
	"log"
	"net"

	"github.com/madsnot/grpc_service/grpc/api"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

			return status.Errorf(codes.ResourceExhausted,
				"%s is rejected because of the limit, please retry later.", info.FullMethod)
		}),

		grpc.UnaryInterceptor(func(ctx context.Context, req interface{},
			info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

			if unaryLimiter.Allow() {
				return handler(ctx, req)
			}

			return nil, status.Errorf(codes.ResourceExhausted,
				"%s is rejected because of the limit, please retry later.", info.FullMethod)
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
		log.Println("Serve", err)
		return
	}
}
