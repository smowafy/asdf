package client

import (
	"crypto/rand"
)

const ClientSecretFileName string = "client-key.rand.asdf"

func readClientSecretKeyFromFile(accountId, filename string) ([]byte, error) {
	return readFromFileName(accountId, filename)
}

func generateAndSaveClientSecret(accountId string) ([]byte, error) {
	if err := statFromFileName(accountId, ClientSecretFileName); err == nil {
		return readClientSecretKeyFromFile(accountId, ClientSecretFileName)
	}

	clientSecret := make([]byte, 16)

	rand.Read(clientSecret)

	if err := writeToFileName(accountId, ClientSecretFileName, clientSecret); err != nil {
		return clientSecret, err
	}

	return clientSecret, nil
}

func (c *AsdfClient) getClientSecret(accountId string) ([]byte, error) {
	if c.clientSecret != nil {
		return c.clientSecret, nil
	}

	return c.createClientSecret(accountId)
}

func (c *AsdfClient) createClientSecret(accountId string) ([]byte, error) {
	var err error

	c.clientSecret, err = generateAndSaveClientSecret(accountId)

	if err != nil {
		return c.clientSecret, err
	}

	return c.clientSecret, nil
}
