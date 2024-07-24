package logger

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger

func InitLogger() *zap.SugaredLogger {
	level, err := zapcore.ParseLevel(viper.GetString("logLevel"))
	if err != nil {
		log.Fatalf("Could not initialize logger")
	}
	env := viper.GetString("env")
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: env == "development",
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:      "json",
		EncoderConfig: encoderConfig,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	logger, _ := config.Build()
	sugar = logger.Sugar()
	return sugar
}

func GetLogger() *zap.SugaredLogger {
	return sugar
}
