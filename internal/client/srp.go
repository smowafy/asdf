package client

import(
	"google.golang.org/grpc"
	"github.com/smowafy/asdf/internal/proto"
	"errors"
	"context"
)

func (asdfClient *AsdfClient) SignUp(masterPassword string, accountId string) error {
	v, err := asdfClient.GetSrpVerifier(masterPassword, accountId)

	if err != nil {
			return err
	}

	id, verifier := v.Encode()

	conn, err := grpc.Dial(":5555", grpc.WithInsecure())

	if err != nil {
		 return err
	}

	defer conn.Close()

	client := proto.NewSrpServerClient(conn)

	res, err := client.SignUp(context.Background(), &proto.ClientVerifier{Id: id, Verif: verifier, AccountId: accountId})

	if err != nil{
		return err
	}

	if !res.Success {
		return errors.New("sign up failed")
	}

	return nil
}

func (asdfClient *AsdfClient) Pake(stream proto.GenericClientStream) ([]byte, error) {
	err := stream.Send(&proto.ClientPayload{Body: []byte(asdfClient.AccountId)})

	if err != nil{
		 return nil, err
	}

	c := asdfClient.srpClient

	creds := c.Credentials()

	err = stream.Send(&proto.ClientPayload{Body: []byte(creds)})

	if err != nil{
		 return nil, err
	}

	serverCreds, err := stream.Recv()

	if err != nil {
		// return nil, err
		panic(err)
	}

	auth, err := c.Generate(string(serverCreds.Body))

	if err != nil {
		 return nil, err
	}

	err = stream.Send(&proto.ClientPayload{Body: []byte(auth)})

	if err != nil {
		return nil, err
	}

	serverProof, err := stream.Recv()

	if err != nil {
		 return nil, err
	}

	if !c.ServerOk(string(serverProof.Body)) {
		return nil, errors.New("server proof not ok")
	}

	return c.RawKey(), nil
}
