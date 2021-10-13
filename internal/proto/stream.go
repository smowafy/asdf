package proto

import (
	grpc "google.golang.org/grpc"
)

type GenericClientStream interface {
	Send(*ClientPayload) error
	Recv() (*ServerPayload, error)
	grpc.ClientStream
}

type GenericServerStream interface {
	Send(*ServerPayload) error
	Recv() (*ClientPayload, error)
	grpc.ServerStream
}
