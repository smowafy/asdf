package server

import(
	"github.com/smowafy/asdf/utils"
	"github.com/smowafy/asdf/internal/proto"
	"fmt"
)

func (asdfServer AsdfServer) GetVault(stream proto.VaultServer_GetVaultServer) error {
	rawKey := asdfServer.rawKey

	encryptedGetVaultPayload, err := stream.Recv()

	if err != nil {
		return err
	}

	vaultId, err := utils.AesDecrypt(rawKey, encryptedGetVaultPayload.Body)

	if err != nil {
		return err
	}

	fmt.Printf("vault Id: %v\n", vaultId)

	vaultBlob, err := asdfServer.db.GetVaultBlob(string(vaultId))

	if err != nil {
		return err
	}

	if vaultBlob == nil {
		err = stream.Send(&proto.ServerPayload{})
	} else {
		fmt.Printf("vault blob: %v\n", vaultBlob)

		encryptedVaultBlob, err := utils.AesEncrypt(rawKey, vaultBlob)

		if err != nil {
			return err
		}

		fmt.Printf("encryptedVaultBlob: %v\n", encryptedVaultBlob)

		err = stream.Send(&proto.ServerPayload{Body: encryptedVaultBlob})
	}

	if err != nil {
		return err
	}

	return nil
}

func (asdfServer AsdfServer) SetVault(stream proto.VaultServer_SetVaultServer) error {
	rawKey := asdfServer.rawKey

	sessionEncryptedVaultId, err := stream.Recv()

	if err != nil {
		return err
	}

	vaultId, err := utils.AesDecrypt(rawKey, sessionEncryptedVaultId.Body)

	if err != nil {
		return err
	}

	fmt.Printf("vault Id: %v\n", vaultId)

	sessionEncryptedVault, err := stream.Recv()

	if err != nil {
		return err
	}

	encryptedVaultBlob, err := utils.AesDecrypt(rawKey, sessionEncryptedVault.Body)

	if err != nil {
		return err
	}

	err = asdfServer.db.SetVaultBlob(string(vaultId), encryptedVaultBlob)

	if err != nil {
		return err
	}

	err = stream.Send(&proto.ServerPayload{})

	return err
}
