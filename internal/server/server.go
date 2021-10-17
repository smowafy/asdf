package server

import (
	"errors"
	"github.com/smowafy/asdf/internal/proto"
	"github.com/smowafy/asdf/internal/server/database"
	"google.golang.org/grpc"
	"log"
	"net"
)

type AsdfServer struct {
	db     database.Database
	rawKey []byte
}

func (s *AsdfServer) Close() error {
	return s.db.Close()
}

func PakeInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if info.FullMethod == "/proto.SrpServer/SignUp" {
		return handler(srv, stream)
	}

	mainSrv, ok := srv.(*AsdfServer)

	if !ok {
		return errors.New("type assertion failed on srv\n")
	}

	rawKey, _, err := mainSrv.Pake(stream)

	if err != nil {
		return err
	}

	workerSrv := AsdfServer{db: mainSrv.db, rawKey: rawKey}

	return handler(workerSrv, stream)
}

func StartServer() {
	log.Printf("listening on :5555\n")

	lis, err := net.Listen("tcp", ":5555")

	if err != nil {
		panic(err)
	}

	db, err := database.NewSqliteDatabase("")

	if err != nil {
		panic(err)
	}

	s := AsdfServer{db: db}

	defer s.Close()

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(PakeInterceptor),
	)

	proto.RegisterSrpServerServer(grpcServer, &s)
	proto.RegisterVaultServerServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
