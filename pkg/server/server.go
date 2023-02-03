package server

import (
	"context"

	"github.com/madsnot/grpc_service/grpc/api"
)

type GRPCServer struct{}

func (s *GRPCServer) SetImage(ctx context.Context, req *api.SetImageRequest) (res *api.SetImageResponse) {
	return res
}

func (s *GRPCServer) GetImagesList(ctx context.Context, req *api.GetImagesListRequest) (res *api.GetImagesListResponse) {
	return res
}

func (s *GRPCServer) GetImage(ctx context.Context, req *api.GetImageRequest) (res *api.GetImageResponse) {
	return res
}
