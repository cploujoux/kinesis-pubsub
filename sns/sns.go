package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	snspackage "github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type SNS struct {
	Client   *snspackage.Client
	TopicARN string
	logger   *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger) (*SNS, error) {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	// Create SNS client
	snsClient := snspackage.NewFromConfig(cfg)
	topicARN := viper.GetString("sns.topic.arn")
	logger.Infow("SNS client started", "topic_arn", topicARN)
	if topicARN == "" {
		return nil, err
	}
	return &SNS{
		Client:   snsClient,
		TopicARN: topicARN,
		logger:   logger,
	}, nil
}
