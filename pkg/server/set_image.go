package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/madsnot/grpc_service/grpc/api"
)

func (s *GRPCServer) SetImage(ctx context.Context, req *api.SetImageRequest) (res *api.SetImageResponse, err error) {
	filePath := req.GetName()

	inputFile, err := os.Open(filePath)
	if err != nil {
		log.Println("Open inputFile")
		return nil, err
	}

	splitPath := strings.Split(filePath, "/")
	fileName := splitPath[len(splitPath)-1]

	newFilePath := fmt.Sprintf("C:/Images/%s", fileName)
	outputFile, err := os.Create(newFilePath)
	if err != nil {
		log.Println("Create outputFile")
		return nil, err
	}

	defer inputFile.Close()
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		log.Println("Write outputFile")
		return nil, err
	}
	return new(api.SetImageResponse), nil
}
