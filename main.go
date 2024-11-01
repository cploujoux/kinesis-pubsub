package main

import (
	"github.com/cploujoux/kinesis-pubsub/api"
	"github.com/cploujoux/kinesis-pubsub/clickhouse"
	"github.com/cploujoux/kinesis-pubsub/config"
	"github.com/cploujoux/kinesis-pubsub/grpc"
	"github.com/cploujoux/kinesis-pubsub/kinesis"
	"github.com/cploujoux/kinesis-pubsub/logger"
	"github.com/cploujoux/kinesis-pubsub/sns"
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
	sns, err := sns.New(sugar)
	if err != nil {
		sugar.Fatalf("Error starting sns client, %v", err)
	}
	clickhouse, err := clickhouse.New(sugar)
	if err != nil {
		sugar.Fatalf("Error starting clickhouse client, %v", err)
	}
	// Start HTTP server
	go func() {
		apiServer := api.New(kinesis, sns, sugar, clickhouse)
		apiServer.Start()
	}()

	// Start gRPC server
	grpcServer := grpc.New(kinesis, sugar)
	grpcServer.Start()
}
