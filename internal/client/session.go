package client

import (
	"encoding/base64"
	"encoding/json"
	"github.com/smowafy/asdf/utils"
	srp "github.com/smowafy/go-srp"
	"os"
)

func SaveSession(c AsdfClient) (string, error) {
	sessionKey := utils.AesGenerateKey()

	encryptedClient, err := encryptedSerialize(c, sessionKey)

	if err != nil {
		return "", err
	}

	if err = os.WriteFile("/tmp/asdf.session", encryptedClient, 0600); err != nil {
		return "", err
	}
	sessionKeyEncoded := make([]byte, base64.StdEncoding.EncodedLen(len(sessionKey)))

	base64.StdEncoding.Encode(sessionKeyEncoded, sessionKey)

	return string(sessionKeyEncoded), nil
}

func LoadSession(sessionKeyEncoded string) (*AsdfClient, error) {
	sessionKey := make([]byte, base64.StdEncoding.DecodedLen(len(sessionKeyEncoded)))

	l, err := base64.StdEncoding.Decode(sessionKey, []byte(sessionKeyEncoded))

	if err != nil {
		return nil, err
	}

	sessionKey = sessionKey[:l]

	encryptedClient, err := os.ReadFile("/tmp/asdf.session")

	if err != nil {
		return nil, err
	}

	return decryptedDeserialize(encryptedClient, sessionKey)
}

func encryptedSerialize(asdfClient AsdfClient, key []byte) ([]byte, error) {
	srpClientKey := utils.AesGenerateKey()

	asdfClient.srpClientKey = srpClientKey

	encSrpClient, err := srp.EncryptedSerialize(asdfClient.srpClient, srpClientKey)

	if err != nil {
		return nil, err
	}

	j, err := json.Marshal(
		struct {
			AccountId    string
			Muk          []byte
			SrpClient    []byte
			SrpClientKey []byte
		}{
			AccountId:    asdfClient.AccountId,
			Muk:          asdfClient.muk,
			SrpClient:    encSrpClient,
			SrpClientKey: asdfClient.srpClientKey,
		})

	if err != nil {
		return nil, err
	}

	encryptedJ, err := utils.AesEncrypt(key, j)

	if err != nil {
		return nil, err
	}

	return encryptedJ, nil
}

func decryptedDeserialize(payload []byte, key []byte) (*AsdfClient, error) {
	j, err := utils.AesDecrypt(key, payload)

	if err != nil {
		return nil, err
	}

	p := &struct {
		AccountId    string
		Muk          []byte
		SrpClient    []byte
		SrpClientKey []byte
	}{}

	err = json.Unmarshal(j, p)

	if err != nil {
		return nil, err
	}

	srpClient, err := srp.DecryptedDeserialize(p.SrpClient, p.SrpClientKey)

	if err != nil {
		return nil, err
	}

	client := &AsdfClient{
		muk:       p.Muk,
		srpClient: srpClient,
		AccountId: p.AccountId,
	}

	if _, err = client.createClientSecret(client.AccountId); err != nil {
		return nil, err
	}

	if _, err := client.createUserKeyPair(client.AccountId, client.muk); err != nil {
		return client, err
	}

	return client, nil
}
