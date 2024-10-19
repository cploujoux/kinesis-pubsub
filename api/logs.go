package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *apiServer) subscribeLogs(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Transfer-Encoding", "chunked")

	a.logger.Info("logs:Subscribed to Clickhouse stream")

	lastTimestamp := time.Time{}

	for {
		query := `
			SELECT *
			FROM moon_v0_otel_logs
			WHERE Timestamp > ?
			ORDER BY Timestamp DESC
			LIMIT 100
		`
		rows, err := a.clickhouse.Db.QueryContext(c.Request.Context(), query, lastTimestamp)
		if err != nil {
			a.logger.Errorf("Error querying logs table: %v", err)
			c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to query logs: %v", err)})
			return
		}

		newLogsFound := false
		for rows.Next() {
			newLogsFound = true
			var log map[string]interface{}
			var timestamp time.Time
			var timestampTime time.Time
			var traceId, spanId, severityText, serviceName, body, resourceSchemaUrl, scopeName, scopeVersion, scopeSchemaUrl string
			var traceFlags, severityNumber uint8
			var resourceAttributes, scopeAttributes, logAttributes map[string]string

			if err := rows.Scan(
				&timestamp,
				&timestampTime,
				&traceId,
				&spanId,
				&traceFlags,
				&severityText,
				&severityNumber,
				&serviceName,
				&body,
				&resourceSchemaUrl,
				&resourceAttributes,
				&scopeSchemaUrl,
				&scopeName,
				&scopeVersion,
				&scopeAttributes,
				&logAttributes,
			); err != nil {
				a.logger.Errorf("Error scanning log: %v", err)
				continue
			}

			log = map[string]interface{}{
				"Timestamp":           timestamp,
				"TimestampTime":       timestampTime,
				"TraceId":             traceId,
				"SpanId":              spanId,
				"TraceFlags":          traceFlags,
				"SeverityText":        severityText,
				"SeverityNumber":      severityNumber,
				"ServiceName":         serviceName,
				"Body":                body,
				"ResourceSchemaUrl":   resourceSchemaUrl,
				"ResourceAttributes":  resourceAttributes,
				"ScopeSchemaUrl":      scopeSchemaUrl,
				"ScopeName":           scopeName,
				"ScopeVersion":        scopeVersion,
				"ScopeAttributes":     scopeAttributes,
				"LogAttributes":       logAttributes,
			}

			a.logger.Infof("Received log in clickhouse: %v", log)
			c.SSEvent("log", log)
			c.Writer.Flush()

			if timestamp.After(lastTimestamp) {
				lastTimestamp = timestamp
			}
		}

		rows.Close()

		if !newLogsFound {
			select {
			case <-c.Request.Context().Done():
				return
			case <-time.After(1 * time.Second): // Poll every second if no new logs
				continue
			}
		}
	}
}
