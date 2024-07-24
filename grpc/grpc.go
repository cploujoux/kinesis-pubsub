package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	"github.com/aws/aws-sdk-go-v2/aws"
	kinesispackage "github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	pb "github.com/cploujoux/kinesis-pubsub/grpc/eventservice"
	"github.com/cploujoux/kinesis-pubsub/kinesis"
	grpcpackage "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type eventServer struct {
	pb.UnimplementedEventServiceServer
	kinesis *kinesis.Kinesis
	logger  *zap.SugaredLogger
}

func (s *eventServer) PushEvent(ctx context.Context, req *pb.PushEventRequest) (*pb.PushEventResponse, error) {
	event := cloudevents.NewEvent()
	event.SetID(req.Id)
	event.SetSource(req.Source)
	event.SetType(req.Type)
	event.SetSpecVersion(req.SpecVersion)
	event.SetData(cloudevents.ApplicationJSON, req.Data)

	eventJSON, err := json.Marshal(event)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to marshal event: %v", err)
	}

	_, err = s.kinesis.Client.PutRecord(ctx, &kinesispackage.PutRecordInput{
		StreamName:   aws.String(s.kinesis.StreamName),
		Data:         eventJSON,
		PartitionKey: aws.String(event.ID()),
	})

	if err != nil {
		s.logger.Errorf("Error pushing event to Kinesis:", err)
		return nil, status.Errorf(codes.Internal, "Failed to push event to Kinesis: %v", err)
	}

	return &pb.PushEventResponse{Message: "Event pushed successfully"}, nil
}

func New(kinesis *kinesis.Kinesis, logger *zap.SugaredLogger) *eventServer {
	return &eventServer{
		kinesis: kinesis,
		logger:  logger,
	}
}

func (s *eventServer) Start() {
	port := viper.GetInt("grpc.port")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		s.logger.Fatalf("Failed to listen: %v", err)
	}
	server := grpcpackage.NewServer()
	pb.RegisterEventServiceServer(server, s)

	s.logger.Infow("gRPC server started", "port", port)
	if err := server.Serve(lis); err != nil {
		s.logger.Fatalf("Failed to serve gRPC: %v", err)
	}
}
