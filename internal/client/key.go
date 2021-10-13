package client

import (
	"crypto/rand"
	"crypto/sha256"
	"io"
	"os"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/crypto/hkdf"
	srp "github.com/opencoff/go-srp"
)

const MukSaltFileName string = "muk-salt.rand.asdf"

func readMukSaltFromFile(filename string) ([]byte, error) {
	return readFromFileName(MukSaltFileName)
}

func generateAndSaveMukSalt() ([]byte, error) {
	if _, err := os.Stat(MukSaltFileName); err == nil {
		return readMukSaltFromFile(MukSaltFileName)
	}

	mukSalt := make([]byte, 8)

	rand.Read(mukSalt)

	if err := writeToFileName(MukSaltFileName, mukSalt); err != nil {
		return mukSalt, err
	}

	return mukSalt, nil
}

func internalKeyGet(clientSecret, salt []byte, masterPassword string, accountID string) ([]byte, error) {
	stretchedMasterPassword, err := scrypt.Key([]byte(masterPassword), salt, 1<<15, 8, 1, 32)

	if err != nil {
		return nil, err
	}

	hkdf := hkdf.New(sha256.New, clientSecret, []byte(accountID), nil)

	intermediateKey := make([]byte, 32)
	_, err = io.ReadFull(hkdf, intermediateKey)

	if err != nil {
		return nil, err
	}

	finalKey := make([]byte, 32)

	for i := range finalKey {
		finalKey[i] = stretchedMasterPassword[i] ^ intermediateKey[i]
	}

	return finalKey, nil
}

func (c *AsdfClient) GetMukSalt() ([]byte, error) {
	if c.mukSalt != nil {
		return c.mukSalt, nil
	}

	return c.createMukSalt()
}

func (c *AsdfClient) createMukSalt() ([]byte, error) {
	var err error

	c.mukSalt, err = generateAndSaveMukSalt()

	if err != nil {
		return c.mukSalt, err
	}

	return c.mukSalt, nil
}


func (c *AsdfClient) GetMuk(masterPassword string, accountId string) ([]byte, error) {
	var err error

	c.muk, err = internalKeyGet(c.clientSecret, c.mukSalt, masterPassword, accountId)

	return c.muk, err
}

func (c *AsdfClient) InitSrpClient(masterPassword string, accountId string) (*srp.Client, error) {
	s, err := srp.New(2048)

	if err != nil {
		return nil, err
	}

	c.srpClient, err = s.NewClient([]byte(accountId), []byte(masterPassword))

	return c.srpClient, err
}

func (c *AsdfClient) GetSrpVerifier(masterPassword string, accountId string) (*srp.Verifier, error) {
	s, err := srp.New(2048)

	if err != nil {
		return nil, err
	}

	return s.Verifier([]byte(accountId), []byte(masterPassword))
}
