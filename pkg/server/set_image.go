package server

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/madsnot/grpc_service/grpc/api"
)

func (s *GRPCServer) SetImage(ctx context.Context, req *api.SetImageRequest) (res *api.SetImageResponse, err error) {
	var newFilePath string

	image := req.GetImage()

	if files, _ := os.ReadDir(servDirPath); files == nil {
		err := os.Mkdir(servDirPath, 0750)
		if err != nil && !os.IsExist(err) {
			return nil, err
		}
	}

	fileNameTemp := fmt.Sprintf("%s/%s-*%s", servDirPath, image.Info.Name, image.Info.Format)
	if files, _ := filepath.Glob(fileNameTemp); len(files) != 0 {
		newFilePath = files[0]
	} else {
		timestamp := time.Now()
		date := strings.Fields(timestamp.String())
		fileName := fmt.Sprintf("%s-%s%s", image.Info.Name, date[0], image.Info.Format)
		newFilePath = fmt.Sprintf("%s/%s", servDirPath, fileName)
	}

	err = os.WriteFile(newFilePath, image.GetData(), 0644)
	if err != nil {
		return nil, err
	}

	return new(api.SetImageResponse), nil
}
