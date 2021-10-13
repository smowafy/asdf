package client

import(
	"os"
	"encoding/json"
	"encoding/base64"
)

func readFromFileName(filename string) ([]byte, error) {
	f, err := os.Open(filename)

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

	if  _, err = f.Read(pEnc); err != nil {
		return nil, err
	}

	l, err := base64.StdEncoding.Decode(payload, pEnc)

	if err != nil {
		return nil, err
	}

	return payload[:l], nil
}

func writeToFileName(filename string, payload []byte) error {
	f, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer f.Close()

	pEnc := make([]byte, base64.StdEncoding.EncodedLen(len(payload)))

	base64.StdEncoding.Encode(pEnc, payload)

	if _, err := f.Write(pEnc); err != nil{
		return err
	}

	return nil
}

func writeToFileNameWithPerm(filename string, payload []byte, perm os.FileMode) error {
	pEnc := make([]byte, base64.StdEncoding.EncodedLen(len(payload)))

	base64.StdEncoding.Encode(pEnc, payload)

	if err := os.WriteFile(filename, pEnc, perm); err != nil{
		return err
	}

	return nil
}


func WriteJsonToFile(filename string, content interface{}) error {
	jsonRes, err := json.Marshal(content)

	if err != nil {
		return err
	}

	return writeToFileNameWithPerm(filename, jsonRes, 0600)
}

func ReadJsonFromFile(filename string, content interface{}) error {
	data, err := readFromFileName(filename)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, content); err != nil {
		return err
	}

	return nil
}
