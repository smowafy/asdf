package database

type Database interface {
	GetVerifier(string, string) (string, error)
	SetVerifier(string, string, string) error
	GetVaultBlob(string) ([]byte, error)
	SetVaultBlob(string, []byte) error
	Close() error
}
