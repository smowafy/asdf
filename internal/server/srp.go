package server

import(
	"github.com/smowafy/asdf/internal/proto"
	srp "github.com/opencoff/go-srp"
	"errors"
	"context"
)

func (asdfServer *AsdfServer) Pake(stream proto.GenericServerStream) ([]byte, string, error) {
	accountIdByte, err := stream.Recv()

	if err != nil {
		 return nil, "", err
	}

	accountId := string(accountIdByte.Body)

	clientIdAndPublicKey, err := stream.Recv()

	if err != nil {
		 return nil, "", err
	}

	id, A, err := srp.ServerBegin(string(clientIdAndPublicKey.Body))

	if err != nil {
		 return nil, "", err
	}

	verifierString, err := asdfServer.db.GetVerifier(id, accountId)

	if err != nil {
		return nil, "", err
	}

	if verifierString == "" {
		return nil, "", errors.New("verifier not found")
	}

	s, v, err := srp.MakeSRPVerifier(verifierString)

	if err != nil {
		return nil, "", err
	}


	srv, err := s.NewServer(v, A)

	serverCredentialsString := srv.Credentials()

	err = stream.Send(&proto.ServerPayload{Body: []byte(serverCredentialsString)})

	if err != nil {
		 return nil, "", err
	}

	clientAuth, err := stream.Recv()

	if err != nil {
		 return nil, "", err
	}

	proof, ok := srv.ClientOk(string(clientAuth.Body))

	if !ok {
		return nil, "", errors.New("client verification not ok")
	}

	err = stream.Send(&proto.ServerPayload{Body: []byte(proof)})

	return srv.RawKey(), accountId, nil
}

func (asdfServer *AsdfServer) SignUp(ctx context.Context, clientVerifier *proto.ClientVerifier) (*proto.VerifierStored, error) {
	err := asdfServer.db.SetVerifier(clientVerifier.Id, clientVerifier.Verif, clientVerifier.AccountId)

	if err != nil {
		return nil, err
	}

	return &proto.VerifierStored{Success: true}, nil
}
