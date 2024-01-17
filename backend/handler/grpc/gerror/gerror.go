package gerror

import (
	"errors"

	"github.com/y-mabuchi/child-memories/backend/domain/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewErr(err error) error {
	var (
		notFound        *model.ErrNotFound
		invalidArgument *model.ErrInvalidArgument
	)

	switch {
	case errors.As(err, &notFound):
		return status.Error(codes.NotFound, notFound.Error())
	case errors.As(err, &invalidArgument):
		return status.Error(codes.InvalidArgument, invalidArgument.Error())
	default:
		return status.Error(codes.Internal, "internal server error")
	}
}
