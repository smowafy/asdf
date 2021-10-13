package server

import(
	"github.com/smowafy/asdf/utils"
	"github.com/smowafy/asdf/internal/proto"
	"fmt"
)

func (asdfServer *AsdfServer) GetVault(stream proto.VaultServer_GetVaultServer) error {
	rawKey, _, err := asdfServer.Pake(stream) // accountId is ignored for now

	if err != nil {
		return err
	}

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
		fmt.Printf("---------\n[server GetVault] this shit is empty\n")
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

func (asdfServer *AsdfServer) SetVault(stream proto.VaultServer_SetVaultServer) error {
	fmt.Printf("[server SetVault] start\n")

	rawKey, _, err := asdfServer.Pake(stream)

	if err != nil {
		return err
	}

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
		fmt.Printf("error is not nil, err: %v\n", err)

		return err
	}

	fmt.Printf("encrypted blob: %v\n", sessionEncryptedVault)

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
