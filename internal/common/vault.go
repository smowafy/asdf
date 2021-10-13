package common

import(
	"github.com/smowafy/asdf/internal/proto"
	protobuf "google.golang.org/protobuf/proto"
)

type Vault map[string][]byte

func EncodeVault(v Vault) ([]byte, error) {
	vaultProto := &proto.Vault{Contents: v}
	return protobuf.Marshal(vaultProto)
}

func DecodeVault(ev []byte) (Vault, error) {
	vaultProto := &proto.Vault{}

	if err := protobuf.Unmarshal(ev, vaultProto); err != nil {
		return nil, err
	}

	return vaultProto.Contents, nil
}
