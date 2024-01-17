package interceptors

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

const (
	realIPKey    = "X-Forwarded-For"
	userAgentKey = "User-Agent"
)

func AccessLoggingInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		startTime := time.Now()

		resp, err := handler(ctx, req)

		c, ok := status.FromError(err)
		if !ok {
			panic(fmt.Sprintf("error is not gPRC status error: %s", err.Error()))
		}

		var realIP string
		var userAgent string

		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if v := md.Get(realIPKey); len(v) > 0 {
				realIP = v[0]
			}

			if v := md.Get(userAgentKey); len(v) > 0 {
				userAgent = v[0]
			}
		}

		logger.Sugar().Infow(
			"access log",
			"method", info.FullMethod,
			"remote_addr", realIP,
			"user_agent", userAgent,
			"request", req,
			"status", c.Code(),
			"response", resp,
			"response_time_second", fmt.Sprintf("%.3f", time.Since(startTime).Seconds()),
		)

		return resp, err
	}
}
