package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/y-mabuchi/child-memories/backend/handler/grpc/interceptors"

	gauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	gzap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	gctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/y-mabuchi/child-memories/backend/domain/utils"
	"github.com/y-mabuchi/child-memories/backend/handler/grpc/generated"
	"github.com/y-mabuchi/child-memories/backend/handler/grpc/services"
	"github.com/y-mabuchi/child-memories/backend/infra/firestore/repository"
	"github.com/y-mabuchi/child-memories/backend/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	client, err := repository.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	regi := registry.NewRegistry(client)

	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	logger, loggerOpt, err := utils.NewZapLogger()
	defer logger.Sync()

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			gctxtags.UnaryServerInterceptor(),
			gzap.UnaryServerInterceptor(logger, loggerOpt),
			interceptors.AccessLoggingInterceptor(logger),
			gauth.UnaryServerInterceptor(interceptors.AuthFunc),
		),
	)

	// Child profile
	childProfileSvc := services.NewChildProfileService(regi)
	generated.RegisterChildProfileServiceServer(grpcServer, childProfileSvc)

	// Family
	familySvc := services.NewFamilyService(regi)
	generated.RegisterFamilyServiceServer(grpcServer, familySvc)

	reflection.Register(grpcServer)

	go func() {
		logger.Info(fmt.Sprintf("start gRPC server port: %v", port))
		grpcServer.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("stopping gRPC server")
	grpcServer.GracefulStop()
}
