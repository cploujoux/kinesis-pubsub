package api

import (
	"fmt"
	"time"

	"github.com/cploujoux/kinesis-pubsub/clickhouse"
	"github.com/cploujoux/kinesis-pubsub/kinesis"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type apiServer struct {
	kinesis   *kinesis.Kinesis
	clickhouse *clickhouse.Clickhouse
	connLimit  chan struct{}
	logger     *zap.SugaredLogger
}

func New(kinesis *kinesis.Kinesis, logger *zap.SugaredLogger, clickhouse *clickhouse.Clickhouse) *apiServer {
	maxConnections := viper.GetInt("api.max.connections")

	return &apiServer{
		kinesis:   kinesis,
		clickhouse: clickhouse,
		connLimit: make(chan struct{}, maxConnections),
		logger:    logger,
	}
}

func (a *apiServer) Start() {
	port := viper.GetInt("api.port")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(ginzap.Ginzap(a.logger.Desugar(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(a.logger.Desugar(), true))

	r.POST("/events", a.pushEvent)
	r.GET("/events", a.subscribeEvents)
	r.GET("/logs", a.subscribeLogs)

	a.logger.Infow("API server started", "port", port, "max_connections", viper.GetInt("api.max.connections"))
	a.logger.Debugw("API server started", "port", port, "max_connections", viper.GetInt("api.max.connections"))
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		a.logger.Fatalf("Failed to run HTTP server: %v", err)
	}
}
