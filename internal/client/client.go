package client

import (
	srp "github.com/smowafy/go-srp"
)

type AsdfClient struct {
	encryptedUserPrivateKey *EncryptedUserPrivateKey
	clientSecret            []byte
	mukSalt                 []byte
	RawKey                  []byte
	muk                     []byte
	srpClient               *srp.Client
	srpClientKey            []byte
	AccountId               string
}

func NewAsdfClient(masterPassword string, accountId string) (*AsdfClient, error) {
	client := &AsdfClient{AccountId: accountId}

	if _, err := client.createClientSecret(accountId); err != nil {
		return client, err
	}

	if _, err := client.createMukSalt(accountId); err != nil {
		return client, err
	}

	muk, err := client.GetMuk(masterPassword, accountId)

	if err != nil {
		return client, err
	}

	if _, err := client.createEncryptedUserPrivateKey(accountId, muk); err != nil {
		return client, err
	}

	if _, err := client.InitSrpClient(masterPassword, accountId); err != nil {
		return client, err
	}

	return client, nil
}
