package database

type MemoryDatabase struct {
	lookup     map[string]string
	vaultBlobs map[string][]byte
}

// check that the interface is implemented
var _ Database = &MemoryDatabase{}

func NewMemoryDatabase() (*MemoryDatabase, error) {
	return &MemoryDatabase{}, nil
}

func (d *MemoryDatabase) GetVerifier(id, accountId string) (string, error) {
	return d.lookup[id], nil
}

func (d *MemoryDatabase) SetVerifier(id, verifier, accountId string) error {
	d.lookup[id] = verifier
	return nil
}

func (d *MemoryDatabase) GetVaultBlob(id string) ([]byte, error) {
	return d.vaultBlobs[id], nil
}

func (d *MemoryDatabase) SetVaultBlob(id string, vault []byte) error {
	d.vaultBlobs[id] = vault
	return nil
}

func (d *MemoryDatabase) Close() error {
	return nil
}
