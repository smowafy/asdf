proto:
	protoc -I './internal/proto/' --go_out=plugins=grpc:./internal internal/proto/vault.proto
	protoc -I './internal/proto/' --go_out=plugins=grpc:./internal internal/proto/srp.proto
