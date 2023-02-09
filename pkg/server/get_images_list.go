package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/madsnot/grpc_service/grpc/api"
	"github.com/madsnot/grpc_service/pkg/errors"
)

func (s *GRPCServer) GetImagesList(ctx context.Context, req *api.GetImagesListRequest) (res *api.GetImagesListResponse, err error) {
	files, _ := os.ReadDir(servDirPath)
	list := make([]string, len(files))
	res = &api.GetImagesListResponse{
		List: list,
	}

	for ind, file := range files {
		info, err := file.Info()
		if err != nil {
			return nil, errors.InternalServerError{Msg: err.Error()}.Error()
		}

		fileFullName := file.Name()

		updateFullDate := info.ModTime()
		updateDate := strings.Fields(updateFullDate.String())[0]

		indLine := strings.Index(fileFullName, "-")
		splitFileFullName := strings.Split(fileFullName[indLine:], ".")
		createDate := splitFileFullName[0][1:]

		fileName := fileFullName[:indLine] + "." + splitFileFullName[len(splitFileFullName)-1]

		imageInfo := fmt.Sprintf("%s | %s | %s", fileName, createDate, updateDate)
		list[ind] = imageInfo
	}

	log.Println("->Get list")

	return res, nil
}
