protoc:
	protoc --go_out=./grpc --go-grpc_out=./grpc ./grpc/event_service.proto

run:
	go run main.go