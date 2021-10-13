package server

import(
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/smowafy/asdf/internal/proto"
	"github.com/smowafy/asdf/internal/server/database"
)

type AsdfServer struct {
	db database.Database
}

func (s *AsdfServer) Close() error {
	return s.db.Close()
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

	grpcServer := grpc.NewServer()

	proto.RegisterSrpServerServer(grpcServer, &s)
	proto.RegisterVaultServerServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
