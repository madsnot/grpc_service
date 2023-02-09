package server

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/madsnot/grpc_service/grpc/api"
	"github.com/madsnot/grpc_service/pkg/errors"
)

func (s *GRPCServer) GetImage(req *api.GetImageRequest, stream api.ImagesHandler_GetImageServer) error {
	fileName := req.GetName()
	fileFormat := req.GetFormat()
	filePathTemp := fmt.Sprintf("%s\\%s-*%s", servDirPath, fileName, fileFormat)

	files, _ := filepath.Glob(filePathTemp)
	if len(files) == 0 {
		return errors.NotExistError{File: fileName + fileFormat}.Error()
	}

	inputFile, err := os.Open(files[0])
	if err != nil {
		return errors.InternalServerError{Msg: err.Error()}.Error()
	}
	defer inputFile.Close()

	info, _ := inputFile.Stat()
	date := info.ModTime().String()
	updateDate := strings.Fields(date)[0]
	filePathTemp = fmt.Sprintf("\\%s-", fileName)
	splitFileName := strings.Split(files[0], filePathTemp)
	indDot := strings.Index(splitFileName[len(splitFileName)-1], ".")
	createDate := splitFileName[len(splitFileName)-1][:indDot]

	res := &api.GetImageResponse{
		Image: &api.Image{
			Info: &api.ImageInfo{
				Name:       fileName,
				Format:     fileFormat,
				CreateDate: createDate,
				UpdateDate: updateDate,
			},
		},
	}

	fileFullName := fileName + fileFormat

	log.Println("->Start download image:", fileFullName)

	chunk := make([]byte, 100000)
	ind := 1
	for {
		n, err := inputFile.Read(chunk)
		if err == io.EOF {
			log.Println("<-End download image:", fileFullName)
			return nil
		}

		res.Image.Data = chunk[:n]
		if err = stream.Send(res); err != nil {
			return errors.InternalServerError{Msg: err.Error()}.Error()
		}

		log.Println(fileFullName, ": Download chunk #", ind)
		ind++
	}

}
