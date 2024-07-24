package main

import (
	"github.com/cploujoux/kinesis-pubsub/api"
	"github.com/cploujoux/kinesis-pubsub/config"
	"github.com/cploujoux/kinesis-pubsub/grpc"
	"github.com/cploujoux/kinesis-pubsub/kinesis"
	"github.com/cploujoux/kinesis-pubsub/logger"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func init() {
	config.InitConfig()
	sugar = logger.InitLogger()
	defer sugar.Sync()
}

func main() {
	kinesis, err := kinesis.New(sugar)
	if err != nil {
		sugar.Fatalf("Error starting kinesis client, %v", err)
	}
	// Start HTTP server
	go func() {
		apiServer := api.New(kinesis, sugar)
		apiServer.Start()
	}()

	// Start gRPC server
	grpcServer := grpc.New(kinesis, sugar)
	grpcServer.Start()
}
