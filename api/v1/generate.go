//go:generate protoc --proto_path=. --proto_path=../../third_party --go_out=paths=source_relative:. --go-gin_out=paths=source_relative:. ./system/*.proto
//go:generate protoc --proto_path=. --proto_path=../../third_party --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./system/*.proto
//go:generate protoc --proto_path=. --proto_path=../../third_party --go_out=paths=source_relative:. --go-http_out=paths=source_relative:. ./system/*.proto
//go:generate protoc --proto_path=. --proto_path=../../third_party --go_out=paths=source_relative:. --go-errors_out=paths=source_relative:. ./system/*.proto
//go:generate protoc --proto_path=. --proto_path=../../third_party --go_out=paths=source_relative:. ./system/*.proto
//go:generate protoc --proto_path=. --proto_path=../../third_party --openapiv2_out . --openapiv2_opt logtostderr=true ./system/*.proto
package v1
