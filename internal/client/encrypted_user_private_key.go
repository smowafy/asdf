package client

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/smowafy/asdf/utils"
	"os"
)

const EncryptedUserKeyFileName string = "encrypted-user-key.rsa.asdf"

type EncryptedUserPrivateKey struct {
	data []byte
}

func NewEncryptedUserPrivateKey(muk []byte) (*EncryptedUserPrivateKey, error) {
	userPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		return nil, err
	}

	fmt.Printf("[NewEncryptedUserPrivateKey] userPrivateKey: %v\n", userPrivateKey)

	encodedUserPrivateKey, err := json.Marshal(userPrivateKey)

	if err != nil {
		return nil, err
	}

	fmt.Printf("[NewEncryptedUserPrivateKey] encodedUserPrivateKey: %v\n", string(encodedUserPrivateKey))

	encryptedUserPrivateKeyPayload, err := utils.AesEncrypt(muk, encodedUserPrivateKey)

	if err != nil {
		return nil, err
	}

	fmt.Printf("[NewEncryptedUserPrivateKey] encryptedUserPrivateKeyPayload: %v\n", encryptedUserPrivateKeyPayload)

	return &EncryptedUserPrivateKey{data: encryptedUserPrivateKeyPayload}, nil
}

func (eupk *EncryptedUserPrivateKey) decryptUserKey(muk []byte) (*rsa.PrivateKey, error) {
	encodedKey, err := utils.AesDecrypt(muk, eupk.data)

	if err != nil {
		return nil, err
	}

	fmt.Printf("[decryptUserKey] encodedKey: %v\n", string(encodedKey))

	key := &rsa.PrivateKey{}

	if err = json.Unmarshal(encodedKey, key); err != nil {
		return nil, err
	}

	fmt.Printf("[decryptUserKey] key: %v\n", key)

	return key, nil
}

func (eupk *EncryptedUserPrivateKey) RsaEncrypt(plaintext []byte, muk []byte) ([]byte, error) {
	rsaKey, err := eupk.decryptUserKey(muk)

	if err != nil {
		return nil, err
	}

	cipher, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &rsaKey.PublicKey, plaintext, nil)

	if err != nil {
		return nil, err
	}

	return cipher, nil
}

func (eupk *EncryptedUserPrivateKey) RsaDecrypt(ciphertext []byte, muk []byte) ([]byte, error) {
	rsaKey, err := eupk.decryptUserKey(muk)

	if err != nil {
		return nil, err
	}

	ev, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, rsaKey, ciphertext, nil)

	if err != nil {
		return nil, err
	}

	return ev, nil
}

func readEncryptedUserPrivateKeyFromFile(filename string) (*EncryptedUserPrivateKey, error) {
	payload, err := readFromFileName(EncryptedUserKeyFileName)

	if err != nil {
		return nil, err
	}

	fmt.Printf("[readEncryptedUserPrivateKeyFromFile] encryptedUserPrivateKeyPayload: %v\n", string(payload))

	return &EncryptedUserPrivateKey{data: payload}, nil
}

func generateAndSaveEncryptedUserPrivateKey(muk []byte) (*EncryptedUserPrivateKey, error) {
	if _, err := os.Stat(EncryptedUserKeyFileName); err == nil {
		return readEncryptedUserPrivateKeyFromFile(EncryptedUserKeyFileName)
	}

	encryptedUserPrivateKey, err := NewEncryptedUserPrivateKey(muk)

	fmt.Printf("[generateAndSaveEncryptedUserPrivateKey] encryptedUserPrivateKey: %v\n", encryptedUserPrivateKey)

	if err != nil {
		return nil, err
	}

	if err = writeToFileName(EncryptedUserKeyFileName, encryptedUserPrivateKey.data); err != nil {
		return nil, err
	}

	return encryptedUserPrivateKey, nil
}

func (c *AsdfClient) createEncryptedUserPrivateKey(muk []byte) (*EncryptedUserPrivateKey, error) {
	var err error

	c.encryptedUserPrivateKey, err = generateAndSaveEncryptedUserPrivateKey(muk)

	if err != nil {
		return c.encryptedUserPrivateKey, err
	}

	return c.encryptedUserPrivateKey, nil
}

func (c *AsdfClient) getEncryptedUserPrivateKey(muk []byte) (*EncryptedUserPrivateKey, error) {
	if c.encryptedUserPrivateKey != nil {
		return c.encryptedUserPrivateKey, nil
	}

	return c.createEncryptedUserPrivateKey(muk)
}
