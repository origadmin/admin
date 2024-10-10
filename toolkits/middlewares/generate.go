package middlewares

//go:generate protoc --proto_path=. --proto_path=../../third_party --go_out=../ --go-http_out=../ --go-grpc_out=../ ./cors/*.proto
