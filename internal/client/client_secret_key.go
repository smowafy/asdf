package client

import(
	"os"
	"crypto/rand"
)

const ClientSecretFileName string = "client-key.rand.asdf"

func readClientSecretKeyFromFile(filename string) ([]byte, error) {
	return readFromFileName(ClientSecretFileName)
}

func generateAndSaveClientSecret() ([]byte, error) {
	if _, err := os.Stat(ClientSecretFileName); err == nil {
		return readClientSecretKeyFromFile(ClientSecretFileName)
	}

	clientSecret := make([]byte, 16)

	rand.Read(clientSecret)

	if err := writeToFileName(ClientSecretFileName, clientSecret); err != nil {
		return clientSecret, err
	}

	return clientSecret, nil
}

func (c *AsdfClient) getClientSecret() ([]byte, error) {
	if c.clientSecret != nil {
		return c.clientSecret, nil
	}

	return c.createClientSecret()
}

func (c *AsdfClient) createClientSecret() ([]byte, error) {
	var err error

	c.clientSecret, err = generateAndSaveClientSecret()

	if err != nil {
		return c.clientSecret, err
	}

	return c.clientSecret, nil
}
