package server

import (
	"log"
	"net"

	"github.com/madsnot/grpc_service/grpc/api"
	"google.golang.org/grpc"
)

var servDirPath = "C:/Images"

type GRPCServer struct {
	api.ImagesHandlerServer
}

func Run() {
	serv := grpc.NewServer()
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
