package server

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/madsnot/grpc_service/grpc/api"
)

func (s *GRPCServer) SetImage(stream api.ImagesHandler_SetImageServer) (err error) {
	var newFilePath string

	if files, _ := os.ReadDir(servDirPath); files == nil {
		err := os.Mkdir(servDirPath, 0750)
		if err != nil && !os.IsExist(err) {
			return err
		}
	}

	req, err := stream.Recv()
	if err == io.EOF {
		return nil
	}

	log.Println("Start upload image:", req.Image.Info.Name+req.Image.Info.Format)

	fileNameTemp := fmt.Sprintf("%s/%s-*%s", servDirPath, req.Image.Info.Name, req.Image.Info.Format)
	if files, _ := filepath.Glob(fileNameTemp); len(files) != 0 {
		newFilePath = files[0]
	} else {
		timestamp := time.Now()
		date := strings.Fields(timestamp.String())
		fileName := fmt.Sprintf("%s-%s%s", req.Image.Info.Name, date[0], req.Image.Info.Format)
		newFilePath = fmt.Sprintf("%s/%s", servDirPath, fileName)
	}

	file, err := os.OpenFile(newFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	fileFullName := req.Image.Info.Name + req.Image.Info.Format
	ind := 1
	_, err = file.Write(req.Image.Data)
	if err != nil {
		return err
	}
	log.Println(fileFullName, ": Upload chunk #", ind)

	for {
		req, err := stream.Recv()
		ind++
		if err == io.EOF {
			log.Println("End upload image:", fileFullName)
			return nil
		}

		_, err = file.Write(req.Image.Data)
		if err != nil {
			return err
		}

		log.Println(fileFullName, ": Upload chunk #", ind)
	}

}
