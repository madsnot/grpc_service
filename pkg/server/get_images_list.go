package server

import (
	"context"

	"github.com/madsnot/grpc_service/grpc/api"
)

func (s *GRPCServer) GetImagesList(ctx context.Context, req *api.GetImagesListRequest) (res *api.GetImagesListResponse, err error) {
	return res, nil
}
