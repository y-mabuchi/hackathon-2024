package utils

import (
	gzap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
)

func NewZapLogger() (*zap.Logger, gzap.Option, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, nil, err
	}

	opt := gzap.WithLevels(
		func(c codes.Code) zapcore.Level {
			var l zapcore.Level
			switch c {
			case codes.OK:
				l = zapcore.InfoLevel
			case codes.NotFound, codes.Unknown:
				l = zapcore.WarnLevel
			case codes.Internal:
				l = zapcore.ErrorLevel
			default:
				l = zapcore.DebugLevel
			}
			return l
		},
	)

	return logger, opt, err
}
