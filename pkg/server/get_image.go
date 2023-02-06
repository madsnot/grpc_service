package server

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/madsnot/grpc_service/grpc/api"
	"github.com/madsnot/grpc_service/pkg/validators"
)

func (s *GRPCServer) GetImage(ctx context.Context, req *api.GetImageRequest) (res *api.GetImageResponse, err error) {
	var splitFileName []string
	servDirPath := "C:/Images"

	fileName := req.GetName()
	filePath := fmt.Sprintf("%s/%s", servDirPath, fileName)

	inputFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	user, err := user.Current()
	if err != nil {
		return nil, err
	}

	userDirPath := fmt.Sprintf("%s\\Downloads", user.HomeDir)
	filePath = fmt.Sprintf("%s\\%s", userDirPath, fileName)

	_, err = os.Open(filePath)
	if err == nil {
		indDot := strings.Index(fileName, ".")
		fileFormat := fileName[indDot:]
		filePathTemp := fmt.Sprintf("%s\\%s *%s", userDirPath, fileName[:indDot], fileFormat)
		files, _ := filepath.Glob(filePathTemp)

		if len(files) == 0 {
			filePath = fmt.Sprintf("%s\\%s 1%s", userDirPath, fileName[:indDot], fileFormat)
		} else {
			splitFileName = strings.Split(files[len(files)-1], ".")
			fieldFileName := strings.Fields(splitFileName[0])

			if validators.ValidNum(fieldFileName[len(fieldFileName)-1]) {
				num, _ := strconv.Atoi(fieldFileName[len(fieldFileName)-1])
				fieldFileName[len(fieldFileName)-1] = " " + strconv.Itoa(num+1)
			}

			filePath = ""
			for _, word := range fieldFileName {
				filePath += word
			}
			filePath += fileFormat
		}
	}

	outputFile, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	defer inputFile.Close()
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return nil, err
	}

	return new(api.GetImageResponse), nil
}
