proto:
	protoc -I './internal/proto/' --go_out=plugins=grpc:./internal internal/proto/vault.proto
	protoc -I './internal/proto/' --go_out=plugins=grpc:./internal internal/proto/srp.proto

server:
	go build -o server ./cmd/server/main.go

client:
	go build -o asdf ./cmd/client/main.go
