# Event Streaming Service

This project implements an event streaming service using AWS Kinesis, supporting both HTTP API and gRPC interfaces. It allows pushing CloudEvents to a Kinesis stream and subscribing to events using server-sent events (SSE).

## Features

- Push CloudEvents to AWS Kinesis stream
- Subscribe to events using Server-Sent Events (SSE)
- Supports both HTTP API and gRPC interfaces

## Prerequisites

- Go 1.16+
- AWS account and credentials configured
- Protocol Buffers compiler (protoc)

## Installation

1. Clone the repository:
git clone https://github.com/cploujoux/kinesis-pubsub.git
cd kinesis-pubsub

2. Install dependencies:
go mod tidy

3. Set up the AWS credentials and configure the Kinesis stream name:
- export AWS_ACCESS_KEY_ID=your_access_key
- export AWS_SECRET_ACCESS_KEY=your_secret_key
- export AWS_REGION=your_region
- export KINESIS_STREAM_NAME=your_stream_name

4. Generate gRPC code:
protoc --go_out=. --go-grpc_out=. event_service.proto

5. Build and run the service:
go build
./kinesis-pubsub

## To improve
- Implement backpressure mechanisms to prevent overwhelming the server or the Kinesis stream.
- Monitor server's resource usage.

## API Documentation

### Push Event (HTTP)

Push a CloudEvent to the Kinesis stream.

- **URL**: `/events`
- **Method**: `POST`
- **Content-Type**: `application/json`

**Request Body**:
```json
{
  "id": "unique-event-id",
  "source": "event-source",
  "type": "event-type",
  "specversion": "1.0",
  "data": {
    "key": "value"
  }
}
```
Response:

Success: HTTP 200
```json
{
  "message": "Event pushed successfully"
}
```

Error: HTTP 400 or 500
```json
{
  "error": "Error message"
}
```

Subscribe to Events (HTTP)
Subscribe to events using Server-Sent Events (SSE).

URL: /events
Method: GET

Response:

Content-Type: text/event-stream
Events are sent as SSE messages

gRPC Documentation
Push Event (gRPC)
Push a CloudEvent to the Kinesis stream using gRPC.
Service: EventService
Method: PushEvent
Request (PushEventRequest):
```protobuf
message PushEventRequest {
  string id = 1;
  string source = 2;
  string type = 3;
  string specversion = 4;
  bytes data = 5;
}
```
Response (PushEventResponse):
```protobuf
message PushEventResponse {
  string message = 1;
}
```

Example gRPC client (Go):
```go
conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
if err != nil {
    log.Fatalf("Failed to connect: %v", err)
}
defer conn.Close()

client := pb.NewEventServiceClient(conn)

resp, err := client.PushEvent(context.Background(), &pb.PushEventRequest{
    Id:          "unique-event-id",
    Source:      "event-source",
    Type:        "event-type",
    Specversion: "1.0",
    Data:        []byte(`{"key": "value"}`),
})

if err != nil {
    log.Fatalf("Failed to push event: %v", err)
}

fmt.Printf("Response: %s\n", resp.Message)
```

## Configuration

The service uses the following environment variables:

- AWS_ACCESS_KEY_ID: AWS access key
- AWS_SECRET_ACCESS_KEY: AWS secret key
- AWS_REGION: AWS region
- KINESIS_STREAM_NAME: Name of the Kinesis stream

Running the Service
The service starts both HTTP and gRPC servers:

- HTTP server runs on port 8080
- gRPC server runs on port 50051

To start the service:
./kinesis-pubsub


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License
This project is licensed under the MIT License.