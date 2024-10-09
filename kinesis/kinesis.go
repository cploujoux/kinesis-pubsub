package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	kinesispackage "github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Kinesis struct {
	Client     *kinesispackage.Client
	StreamName string
	logger     *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger) (*Kinesis, error) {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	// Create Kinesis client
	kinesisClient := kinesispackage.NewFromConfig(cfg)
	streamName := viper.GetString("kinesis.stream.name")
	logger.Infow("Kinesis client started", "stream_name", streamName)
	if streamName == "" {
		return nil, err
	}
	return &Kinesis{
		Client:     kinesisClient,
		StreamName: streamName,
		logger:     logger,
	}, nil
}
