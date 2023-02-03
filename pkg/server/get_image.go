package server

import (
	"context"

	"github.com/madsnot/grpc_service/grpc/api"
)

func (s *GRPCServer) GetImage(ctx context.Context, req *api.GetImageRequest) (res *api.GetImageResponse, err error) {
	return res, nil
}
