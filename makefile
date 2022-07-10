proto:
	protoc -I=./protofiles  --go_out=./src/api/grpc/ptypes  --go_opt=paths=source_relative --go-grpc_out=./src/api/grpc/ptypes --go-grpc_opt=paths=source_relative  ./protofiles/*.proto
