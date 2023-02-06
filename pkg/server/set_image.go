package server

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/madsnot/grpc_service/grpc/api"
)

func (s *GRPCServer) SetImage(ctx context.Context, req *api.SetImageRequest) (res *api.SetImageResponse, err error) {
	filePath := req.GetName()

	inputFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	splitPath := strings.Split(filePath, "/")
	fileName := splitPath[len(splitPath)-1]

	dirPath := "C:/Images"
	if filesList, _ := os.ReadDir(dirPath); filesList == nil {
		err := os.Mkdir(dirPath, 0750)
		if err != nil && !os.IsExist(err) {
			return nil, err
		}
	}

	newFilePath := fmt.Sprintf("%s/%s", dirPath, fileName)

	outputFile, err := os.Create(newFilePath)
	if err != nil {
		return nil, err
	}

	defer inputFile.Close()
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return nil, err
	}

	return new(api.SetImageResponse), nil
}
