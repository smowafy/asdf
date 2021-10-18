package client

import (
	"encoding/base64"
	"fmt"
	"os"
)

func fullFilename(dirname, filename string) string {
	return fmt.Sprintf("%s/%s", dirname, filename)
}

func statFromFileName(dirname, filename string) error {
	_, err := os.Stat(fullFilename(dirname, filename))

	return err
}

func readFromFileName(dirname, filename string) ([]byte, error) {
	fullFilename := fullFilename(dirname, filename)

	f, err := os.Open(fullFilename)

	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()

	if err != nil {
		return nil, err
	}

	encSize := int(stat.Size())

	payload := make([]byte, base64.StdEncoding.DecodedLen(encSize))
	pEnc := make([]byte, encSize)

	if _, err = f.Read(pEnc); err != nil {
		return nil, err
	}

	l, err := base64.StdEncoding.Decode(payload, pEnc)

	if err != nil {
		return nil, err
	}

	return payload[:l], nil
}

func writeToFileName(dirname, filename string, payload []byte) error {
	err := os.MkdirAll(dirname, 0700)

	if err != nil {
		return err
	}

	fullFilename := fullFilename(dirname, filename)

	f, err := os.Create(fullFilename)

	if err != nil {
		return err
	}

	defer f.Close()

	pEnc := make([]byte, base64.StdEncoding.EncodedLen(len(payload)))

	base64.StdEncoding.Encode(pEnc, payload)

	if _, err := f.Write(pEnc); err != nil {
		return err
	}

	return nil
}

func writeToFileNameWithPerm(filename string, payload []byte, perm os.FileMode) error {
	pEnc := make([]byte, base64.StdEncoding.EncodedLen(len(payload)))

	base64.StdEncoding.Encode(pEnc, payload)

	if err := os.WriteFile(filename, pEnc, perm); err != nil {
		return err
	}

	return nil
}
