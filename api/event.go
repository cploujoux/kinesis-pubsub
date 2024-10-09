package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/gin-gonic/gin"
)

func (a *apiServer) pushEvent(c *gin.Context) {
	event, err := cloudevents.NewEventFromHTTPRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid CloudEvent format:%v", err)})
		return
	}

	eventJSON, err := json.Marshal(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal event"})
		return
	}

	_, err = a.kinesis.Client.PutRecord(context.TODO(), &kinesis.PutRecordInput{
		StreamName:   aws.String(a.kinesis.StreamName),
		Data:         eventJSON,
		PartitionKey: aws.String(event.ID()),
	})

	if err != nil {
		a.logger.Error("Error pushing event to Kinesis:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to push event to Kinesis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event pushed successfully"})
}

// subscribeEvents handles the GET /events endpoint to subscribe to CloudEvents from Kinesis
func (a *apiServer) subscribeEvents(c *gin.Context) {
	select {
	case a.connLimit <- struct{}{}:
		defer func() { <-a.connLimit }()
	default:
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Maximum connections reached"})
		return
	}

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Transfer-Encoding", "chunked")

	shardIterators, err := a.getShardIterators()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get shard iterators"})
		return
	}

	a.logger.Info("event:Subscribed to Kinesis stream")
	for {
		select {
		case <-c.Request.Context().Done():
			return
		default:
			for shardId, iterator := range shardIterators {
				output, err := a.kinesis.Client.GetRecords(context.TODO(), &kinesis.GetRecordsInput{
					ShardIterator: iterator,
				})

				if err != nil {
					a.logger.Errorf("Error getting records from shard %s: %v", shardId, err)
					continue
				}

				for _, record := range output.Records {
					var event cloudevents.Event
					if err := json.Unmarshal(record.Data, &event); err != nil {
						a.logger.Errorf("Error unmarshalling event: %v", err)
						continue
					}

					c.SSEvent("message", event)
					c.Writer.Flush()
				}

				shardIterators[shardId] = output.NextShardIterator
			}
			time.Sleep(time.Second)
		}
	}
}

// getShardIterators retrieves shard iterators for all shards in the Kinesis stream.
func (a *apiServer) getShardIterators() (map[string]*string, error) {
	shardIterators := make(map[string]*string)

	// List all shards
	var exclusiveStartShardId *string
	for {
		shards, err := a.kinesis.Client.ListShards(context.TODO(), &kinesis.ListShardsInput{
			StreamName:            aws.String(a.kinesis.StreamName),
			ExclusiveStartShardId: exclusiveStartShardId,
		})

		if err != nil {
			return nil, err
		}

		for _, shard := range shards.Shards {
			shardIterator, err := a.kinesis.Client.GetShardIterator(context.TODO(), &kinesis.GetShardIteratorInput{
				StreamName:        aws.String(a.kinesis.StreamName),
				ShardId:           shard.ShardId,
				ShardIteratorType: types.ShardIteratorTypeLatest,
			})

			if err != nil {
				return nil, err
			}

			shardIterators[*shard.ShardId] = shardIterator.ShardIterator
		}

		if shards.NextToken == nil {
			break
		}
		exclusiveStartShardId = shards.Shards[len(shards.Shards)-1].ShardId
	}

	if len(shardIterators) == 0 {
		return nil, fmt.Errorf("no shards found in the stream")
	}

	return shardIterators, nil
}
