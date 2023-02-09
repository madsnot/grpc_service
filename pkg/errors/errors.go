package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ConnLimitExceededError struct {
	Method string
}

func (err ConnLimitExceededError) Error() error {
	return status.Errorf(codes.NotFound, fmt.Sprintf("%s is rejected because of the limit, please retry later.", err.Method))
}

type NotExistError struct {
	File string
}

func (err NotExistError) Error() error {
	return status.Errorf(codes.NotFound, fmt.Sprintf("%s file doesn`t exist.", err.File))
}

type InvalidNameError struct {
	File string
}

func (err InvalidNameError) Error() error {
	return status.Errorf(codes.InvalidArgument, fmt.Sprintf("%s file has invalid name.", err.File))
}

type InternalServerError struct {
	Msg string
}

func (err InternalServerError) Error() error {
	return status.Errorf(codes.Internal, err.Msg)
}
