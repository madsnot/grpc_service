package server

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/madsnot/grpc_service/grpc/api"
)

func (s *GRPCServer) GetImage(req *api.GetImageRequest, stream api.ImagesHandler_GetImageServer) (err error) {
	fileName := req.GetName()
	fileFormat := req.GetFormat()
	filePathTemp := fmt.Sprintf("%s\\%s-*%s", servDirPath, fileName, fileFormat)

	files, _ := filepath.Glob(filePathTemp)
	if len(files) == 0 {
		return fmt.Errorf("this file doesn`t exist")
	}

	inputFile, err := os.Open(files[0])
	if err != nil {
		return err
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

	log.Println("Start download image:", fileName+fileFormat)

	chunk := make([]byte, 100000)
	ind := 1
	for {
		n, err := inputFile.Read(chunk)
		if err == io.EOF {
			log.Println("End download image:", fileName+fileFormat)
			return nil
		}

		res.Image.Data = chunk[:n]
		if err = stream.Send(res); err != nil {
			return err
		}

		log.Println(fileName+fileFormat, ": Download chunk #", ind)
		ind++
	}

}
