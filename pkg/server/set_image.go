package server

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/madsnot/grpc_service/grpc/api"
)

func (s *GRPCServer) SetImage(ctx context.Context, req *api.SetImageRequest) (res *api.SetImageResponse, err error) {
	image := req.GetImage()

	dirPath := "C:/Images"
	if filesList, _ := os.ReadDir(dirPath); filesList == nil {
		err := os.Mkdir(dirPath, 0750)
		if err != nil && !os.IsExist(err) {
			return nil, err
		}
	}

	timestamp := time.Now()
	date := strings.Fields(timestamp.String())
	fileName := fmt.Sprintf("%s-%s%s", image.Info.Name, date[0], image.Info.Format)
	newFilePath := fmt.Sprintf("%s/%s", dirPath, fileName)

	err = os.WriteFile(newFilePath, image.GetData(), 0644)
	if err != nil {
		return nil, err
	}

	return new(api.SetImageResponse), nil
}
