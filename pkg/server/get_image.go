package server

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/madsnot/grpc_service/grpc/api"
)

func (s *GRPCServer) GetImage(ctx context.Context, req *api.GetImageRequest) (res *api.GetImageResponse, err error) {
	fileName := req.GetName()
	fileFormat := req.GetFormat()
	filePathTemp := fmt.Sprintf("%s\\%s-*%s", servDirPath, fileName, fileFormat)

	files, _ := filepath.Glob(filePathTemp)
	if len(files) == 0 {
		return nil, fmt.Errorf("this file doesn`t exist")
	}

	inputFile, err := os.Open(files[0])
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	info, err := inputFile.Stat()
	size := info.Size()
	data := make([]byte, size)
	date := info.ModTime().String()
	updateDate := strings.Fields(date)[0]
	filePathTemp = fmt.Sprintf("\\%s-", fileName)
	splitFileName := strings.Split(files[0], filePathTemp)
	indDot := strings.Index(splitFileName[len(splitFileName)-1], ".")
	createDate := splitFileName[len(splitFileName)-1][:indDot]

	res = &api.GetImageResponse{
		Image: &api.Image{
			Info: &api.ImageInfo{
				Name:       fileName,
				Format:     fileFormat,
				CreateDate: createDate,
				UpdateDate: updateDate,
			},
			Data: data,
		},
	}

	_, err = inputFile.Read(res.Image.Data)
	if err != nil {
		return nil, err
	}

	return res, nil
}
