package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/smowafy/asdf/internal/common"
	"github.com/smowafy/asdf/internal/proto"
	"github.com/smowafy/asdf/utils"
	"google.golang.org/grpc"
)

func (c *AsdfClient) EncryptVault(v common.Vault) ([]byte, error) {
	ev, err := common.EncodeVault(v)

	if err != nil {
		return nil, err
	}

	// TODO: add account ID as label

	privKey, err := c.getEncryptedUserPrivateKey(c.muk)

	if err != nil {
		return nil, err
	}

	cipher, err := privKey.RsaEncrypt(ev, c.muk)

	if err != nil {
		return nil, err
	}

	return cipher, nil
}

func (c *AsdfClient) DecryptVault(cipher []byte) (common.Vault, error) {
	privKey, err := c.getEncryptedUserPrivateKey(c.muk)

	if err != nil {
		return nil, err
	}

	ev, err := privKey.RsaDecrypt(cipher, c.muk)

	if err != nil {
		return nil, err
	}

	return common.DecodeVault(ev)
}

func (asdfClient *AsdfClient) GetVault() (common.Vault, error) {
	conn, err := grpc.Dial(":5555", grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := proto.NewVaultServerClient(conn)

	stream, err := client.GetVault(context.Background())

	if err != nil {
		return nil, err
	}

	key, err := asdfClient.Pake(stream)

	if err != nil {
		return nil, err
	}

	encVaultId, err := utils.AesEncrypt(key, []byte(asdfClient.AccountId))

	if err != nil {
		return nil, err
	}

	err = stream.Send(&proto.ClientPayload{Body: encVaultId})

	if err != nil {
		return nil, err
	}

	encryptedVaultResponse, err := stream.Recv()

	if err != nil {
		return nil, err
	}

	sessionEncryptedVaultBlob := encryptedVaultResponse.Body

	fmt.Printf("---------\n\n[GetVault] sessionEncryptedVaultBlob: %v, nil? = %v\n", sessionEncryptedVaultBlob)

	if sessionEncryptedVaultBlob == nil {
		return nil, nil
	}

	encryptedVaultBlob, err := utils.AesDecrypt(key, sessionEncryptedVaultBlob)

	if err != nil {
		return nil, err
	}

	vault, err := asdfClient.DecryptVault(encryptedVaultBlob)

	if err != nil {
		return nil, err
	}

	return vault, nil
}

func (asdfClient *AsdfClient) SetVault(vault common.Vault) error {
	conn, err := grpc.Dial(":5555", grpc.WithInsecure())

	if err != nil {
		return err
	}

	defer conn.Close()

	client := proto.NewVaultServerClient(conn)

	stream, err := client.SetVault(context.Background())

	if err != nil {
		return err
	}

	key, err := asdfClient.Pake(stream)

	if err != nil {
		return err
	}

	encVaultId, err := utils.AesEncrypt(key, []byte(asdfClient.AccountId))

	if err != nil {
		return err
	}

	err = stream.Send(&proto.ClientPayload{Body: encVaultId})

	if err != nil {
		return err
	}

	encryptedVaultBlob, err := asdfClient.EncryptVault(vault)

	if err != nil {
		return err
	}

	sessionEncryptedVaultBlob, err := utils.AesEncrypt(key, encryptedVaultBlob)

	err = stream.Send(&proto.ClientPayload{Body: sessionEncryptedVaultBlob})

	if err != nil {
		return err
	}

	ack, err := stream.Recv()

	if err != nil {
		return err
	}

	if ack.Body != nil {
		return errors.New(string(ack.Body))
	}

	return nil
}

func (c *AsdfClient) GetItem(itemKey string) ([]byte, error) {
	vault, err := c.GetVault()

	if err != nil {
		return nil, err
	}

	if vault == nil {
		return nil, nil
	}

	return vault[itemKey], nil
}

func (c *AsdfClient) SetItem(itemKey string, value []byte) error {
	// TODO: Get vault, set key in vault then send vault back, also extract Pake
	// from all this shit
	vault, err := c.GetVault()

	if err != nil {
		return err
	}

	if vault == nil {
		vault = make(common.Vault)
	}

	vault[itemKey] = value

	err = c.SetVault(vault)

	if err != nil {
		return err
	}

	return nil
}
