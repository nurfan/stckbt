package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stockbit-service-movie/util/errors"
)

var grpcCode = map[int16]codes.Code{
	400: codes.InvalidArgument,
	200: codes.OK,
	201: codes.OK,
}


	func errorMapping(err *errors.UniError) error {
	return status.Error(grpcCode[err.Code], err.Message)
}
