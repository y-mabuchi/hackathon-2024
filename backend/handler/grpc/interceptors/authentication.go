package interceptors

import (
	"context"

	gauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const AuthKey = "crowdlog"

func AuthFunc(ctx context.Context) (context.Context, error) {
	key, err := gauth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, status.Errorf(
			codes.Unauthenticated,
			"could not read auth token: %v",
			err,
		)
	}

	if key != AuthKey {
		return nil, status.Errorf(
			codes.Unauthenticated,
			"invalid api password: %v",
			err,
		)
	}

	newCtx := context.WithValue(ctx, "result", "ok")

	return newCtx, nil
}
