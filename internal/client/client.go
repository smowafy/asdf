package client

import(
	srp "github.com/opencoff/go-srp"
	"encoding/json"
)

type AsdfClient struct {
	encryptedUserPrivateKey *EncryptedUserPrivateKey
	clientSecret []byte
	mukSalt []byte
	RawKey []byte
	muk []byte
	srpClient *srp.Client
	AccountId string
}

func NewAsdfClient(masterPassword string, accountId string) (*AsdfClient, error) {
	client := &AsdfClient{AccountId: accountId}

	if _, err := client.createClientSecret(); err != nil {
		return client, err
	}

	if _, err := client.createMukSalt(); err != nil {
		return client, err
	}

	muk, err := client.GetMuk(masterPassword, accountId)

	if err != nil {
		return client, err
	}

	if _, err := client.createEncryptedUserPrivateKey(muk); err != nil {
		return client, err
	}

	if _, err := client.InitSrpClient(masterPassword, accountId); err != nil{
		return client, err
	}

	return client, nil
}

func (asdfClient AsdfClient) MarshalJSON() ([]byte, error) {
	encPrivKey, err := json.Marshal(asdfClient.encryptedUserPrivateKey)

	if err != nil{
		return nil, err
	}

	encSrpClient, err := json.Marshal(asdfClient.srpClient)

	if err != nil{
		return nil, err
	}

	j, err := json.Marshal(
		struct {
			EncryptedUserPrivateKey string
			ClientSecret []byte
			MukSalt []byte
			RawKey []byte
			Muk []byte
			SrpClient string
		}{
			EncryptedUserPrivateKey: string(encPrivKey),
			ClientSecret: asdfClient.clientSecret,
			MukSalt: asdfClient.mukSalt,
			RawKey: asdfClient.RawKey,
			Muk: asdfClient.muk,
			SrpClient: string(encSrpClient),
		})

	if err != nil {
		return nil, err
	}

	return j, nil
}