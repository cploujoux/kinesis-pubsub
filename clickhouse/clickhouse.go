package clickhouse

import (
	"crypto/tls"
	"database/sql"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Clickhouse struct {
	Db *sql.DB
}

func New(logger *zap.SugaredLogger) (*Clickhouse, error) {
	clickhouseDb := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", viper.GetString("clickhouse.host"), viper.GetInt("clickhouse.port"))},
		Auth: clickhouse.Auth{
			Database: viper.GetString("clickhouse.database"),
			Username: viper.GetString("clickhouse.user"),
			Password: viper.GetString("clickhouse.password"),
		},
		ClientInfo: clickhouse.ClientInfo{
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "pub-sub-go-client", Version: "0.1"},
			},
		},
		TLS: &tls.Config{},
	})
	return &Clickhouse{
		Db: clickhouseDb,
	}, nil
}
