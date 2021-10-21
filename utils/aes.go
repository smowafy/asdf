package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	//	"fmt"
)

func AesEncrypt(key []byte, payload []byte) ([]byte, error) {
	//	fmt.Printf("------------- AES encrypt begin ----------------\n")
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)

	nonce := make([]byte, gcm.NonceSize())

	//	fmt.Printf("nonce size: %v\n", gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	//	fmt.Printf("nonce missing?: %v\n", nonce)

	rawCipherText := gcm.Seal(nonce, nonce, payload, nil)

	//	fmt.Printf("rawCipherText: %v, len = %v\n", rawCipherText, len(rawCipherText))

	encoding := base64.StdEncoding

	encodedLen := encoding.EncodedLen(len(rawCipherText))

	encodedCipherText := make([]byte, encodedLen)

	encoding.Encode(encodedCipherText, rawCipherText)

	//	fmt.Printf("encodedCipherText: %v\n", encodedCipherText)

	//	fmt.Printf("------------- AES encrypt end ----------------\n")

	return encodedCipherText, nil
}

func AesDecrypt(key []byte, encodedCipherText []byte) ([]byte, error) {
	//	fmt.Printf("------------- AES decrypt begin ----------------\n")

	encoding := base64.StdEncoding

	//	fmt.Printf("encodedCipherText: %v\n", encodedCipherText)

	decodedLen := encoding.DecodedLen(len(encodedCipherText))

	//	fmt.Printf("decodedLen: %v\n", decodedLen)

	nonceAndCipherText := make([]byte, decodedLen)

	actual, err := encoding.Decode(nonceAndCipherText, encodedCipherText)

	if err != nil {
		return nil, err
	}

	nonceAndCipherText = nonceAndCipherText[:actual]

	//	fmt.Printf("nonceAndCipherText: %v\n", nonceAndCipherText)

	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()

	//	fmt.Printf("nonce size: %v\n", nonceSize)

	nonce, cipherText := nonceAndCipherText[:nonceSize], nonceAndCipherText[nonceSize:]

	//	fmt.Printf("nonce: %v\n cipherText: %v\n", nonce, cipherText)

	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)

	if err != nil {
		return nil, err
	}

	//	fmt.Printf("------------- AES decrypt end ----------------\n")

	return plaintext, nil
}
