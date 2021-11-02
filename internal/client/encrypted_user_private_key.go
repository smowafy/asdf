package client

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/smowafy/asdf/utils"
	"golang.org/x/crypto/nacl/box"
	"io"
)

const EncryptedUserKeyFileName string = "encrypted-user-key.rsa.asdf"

const EncryptedUserPrivateKeyDefaultFileName string = "enc-priv.25519.asdf"
const UserPublicKeyDefaultFileName string = "pub.25519.asdf"

type EncryptedUserPrivateKey []byte
type UserPublicKeyPtr *[32]byte

type UserKeyPair struct {
	encPrivateKey EncryptedUserPrivateKey
	publicKey     UserPublicKeyPtr
}

func NewUserKeyPair(muk []byte) (*UserKeyPair, error) {
	userPrivateKey, userPublicKey, err := box.GenerateKey(rand.Reader)

	if err != nil {
		return nil, err
	}

	fmt.Printf("[NewEncryptedUserPrivateKey] key: %v\n", userPrivateKey)

	encodedUserPrivateKey, err := json.Marshal(userPrivateKey)

	if err != nil {
		return nil, err
	}

	encryptedUserPrivateKeyPayload, err := utils.AesEncrypt(muk, encodedUserPrivateKey)

	if err != nil {
		return nil, err
	}

	return &UserKeyPair{
		encPrivateKey: encryptedUserPrivateKeyPayload,
		publicKey:     userPublicKey,
	}, nil
}

func (eupk *EncryptedUserPrivateKey) decryptUserKey(muk []byte) (*[32]byte, error) {
	encodedKey, err := utils.AesDecrypt(muk, *eupk)

	if err != nil {
		return nil, err
	}

	fmt.Printf("[decryptUserKey] encodedKey: %v\n", string(encodedKey))

	var key [32]byte

	if err = json.Unmarshal(encodedKey, &key); err != nil {
		return nil, err
	}

	fmt.Printf("[decryptUserKey] key: %v\n", key)

	return &key, nil
}

func (ukp *UserKeyPair) SelfEncrypt(plaintext []byte, muk []byte) ([]byte, error) {
	key, err := ukp.encPrivateKey.decryptUserKey(muk)

	if err != nil {
		return nil, err
	}

	var nonce [24]byte

	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return nil, err
	}

	cipher := box.Seal(nonce[:], plaintext, &nonce, ukp.publicKey, key)

	return cipher, nil
}

func (ukp *UserKeyPair) SelfDecrypt(ciphertext []byte, muk []byte) ([]byte, error) {
	key, err := ukp.encPrivateKey.decryptUserKey(muk)

	if err != nil {
		return nil, err
	}

	var nonce [24]byte

	copy(nonce[:], ciphertext[:24])

	ev, ok := box.Open(nil, ciphertext[24:], &nonce, ukp.publicKey, key)

	if !ok {
		return nil, errors.New("decryption failure")
	}

	return ev, nil
}

func readEncryptedUserPrivateKeyFromFile(accountId string) (EncryptedUserPrivateKey, error) {
	return readFromFileName(accountId, EncryptedUserPrivateKeyDefaultFileName)
}

func readUserPublicKeyFromFile(accountId string) (UserPublicKeyPtr, error) {
	payload, err := readFromFileName(accountId, UserPublicKeyDefaultFileName)

	if err != nil {
		return nil, err
	}

	var res [32]byte

	copy(res[:], payload)

	return &res, nil
}

func readUserKeyPairFromFiles(accountId string) (*UserKeyPair, error) {
	priv, err := readEncryptedUserPrivateKeyFromFile(accountId)

	if err != nil {
		return nil, err
	}

	pub, err := readUserPublicKeyFromFile(accountId)

	if err != nil {
		return nil, err
	}

	return &UserKeyPair{
		encPrivateKey: priv,
		publicKey:     pub,
	}, nil
}

func keyPairExists(accountId string) bool {
	if err := statFromFileName(accountId, EncryptedUserPrivateKeyDefaultFileName); err != nil {
		return false
	}
	if err := statFromFileName(accountId, UserPublicKeyDefaultFileName); err != nil {
		return false
	}

	return true
}

func generateAndSaveUserKeyPair(accountId string, muk []byte) (*UserKeyPair, error) {
	if keyPairExists(accountId) {
		return readUserKeyPairFromFiles(accountId)
	}

	userKeyPair, err := NewUserKeyPair(muk)

	if err != nil {
		return nil, err
	}

	if err = writeToFileName(
		accountId,
		EncryptedUserPrivateKeyDefaultFileName,
		userKeyPair.encPrivateKey,
	); err != nil {
		return nil, err
	}

	if err = writeToFileName(
		accountId,
		UserPublicKeyDefaultFileName,
		(*userKeyPair.publicKey)[:],
	); err != nil {
		return nil, err
	}

	return userKeyPair, nil
}

func (c *AsdfClient) createUserKeyPair(accountId string, muk []byte) (*UserKeyPair, error) {
	var err error

	c.userKeyPair, err = generateAndSaveUserKeyPair(accountId, muk)

	if err != nil {
		return nil, err
	}

	return c.userKeyPair, nil
}

func (c *AsdfClient) getUserKeyPair(accountId string, muk []byte) (*UserKeyPair, error) {
	if c.userKeyPair != nil {
		return c.userKeyPair, nil
	}

	return c.createUserKeyPair(accountId, muk)
}
