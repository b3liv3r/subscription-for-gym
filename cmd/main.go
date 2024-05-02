package main

import (
	loggerx "github.com/b3liv3r/logger"
	subv1 "github.com/b3liv3r/protos-for-gym/gen/go/subscription"
	"github.com/b3liv3r/subscription-for-gym/config"
	"github.com/b3liv3r/subscription-for-gym/modules/db"
	"github.com/b3liv3r/subscription-for-gym/modules/subscription/repository"
	"github.com/b3liv3r/subscription-for-gym/modules/subscription/service"
	server "github.com/b3liv3r/subscription-for-gym/modules/subscription/srpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	appConf := config.MustLoadConfig(".env")

	logger := loggerx.InitLogger(appConf.Name, appConf.Production)

	sqlDB, err := db.NewSqlDB(logger, appConf.Db)
	if err != nil {
		logger.Fatal("failed to connect to db", zap.Error(err))
	}

	repo := repository.NewSubscriptionRepositoryDB(sqlDB)
	service := service.NewSubscriptionService(repo, logger)
	s := InitRPC(service)
	lis, err := net.Listen("tcp", appConf.GrpcServerPort)
	if err != nil {
		logger.Error("failed to listen:", zap.Error(err))
	}
	logger.Info("grpc server listening at", zap.Stringer("address", lis.Addr()))
	if err = s.Serve(lis); err != nil {
		logger.Fatal("failed to serve:", zap.Error(err))
	}
}

func InitRPC(wservice service.Subscriber) *grpc.Server {
	s := grpc.NewServer()
	subv1.RegisterSubscriptionServer(s, server.NewSubscriptionRPCServer(wservice))

	return s
}
