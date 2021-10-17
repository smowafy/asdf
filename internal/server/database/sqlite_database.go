package database

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

const (
	DefaultConnStr = "./asdf.db"

	// TODO: think about schema (account & verifier IDs)
	createVerifierTableStmt = `
	CREATE TABLE IF NOT EXISTS verifiers (
		id INTEGER PRIMARY KEY,
		verifier_id TEXT,
		verifier TEXT,
		account_id TEXT
	)
	`
	createVerifierIndexStmt=`
	CREATE UNIQUE INDEX IF NOT EXISTS verifiers_on_verifier_id
	ON verifiers(verifier_id)
	`
	createVaultTableStmt = `
	CREATE TABLE IF NOT EXISTS vaults (
		id INTEGER PRIMARY KEY,
		vault_id TEXT,
		blob BLOB
	)
	`
	createVaultIndexStmt = `
	CREATE UNIQUE INDEX IF NOT EXISTS vaults_on_vault_id
	ON vaults(vault_id)
	`

	queryVerifierStmt = `
	SELECT verifier FROM verifiers WHERE
	verifier_id = ? AND account_id = ?
	`

	SetVerifierStmt = `
	INSERT INTO verifiers(verifier_id, verifier, account_id) VALUES (?, ?, ?)
	`

	queryVaultStmt = `
	SELECT blob FROM vaults WHERE
	vault_id = ?
	`

	insertOrUpdateVaultStmt = `
	INSERT INTO vaults(vault_id, blob) VALUES (?, ?)
	ON CONFLICT(vault_id) DO
	UPDATE SET blob = ?
	`
)

type SqliteDatabase struct {
	db *sql.DB
}

// check that the interface is implemented
var _ Database = &SqliteDatabase{}

func NewSqliteDatabase(connStr string) (*SqliteDatabase, error) {
	if connStr == "" {
		connStr = DefaultConnStr
	}

	db, err := sql.Open("sqlite3", connStr)

	if err != nil {
		return nil, err
	}

	d := &SqliteDatabase{db: db}

	if err = d.bootstrapSchema(); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *SqliteDatabase) GetVerifier(id, accountId string) (string, error) {
	var verifier string

	if err := d.db.QueryRow(queryVerifierStmt, id, accountId).Scan(&verifier); err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}

		return "", err
	}

	return verifier, nil
}

func (d *SqliteDatabase) SetVerifier(id, verifier, accountId string) error {
	res, err := d.db.Exec(SetVerifierStmt, id, verifier, accountId)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	lastid, err := res.LastInsertId()

	if err != nil {
		return err
	}

	fmt.Printf("[db set verifier] rows affected: %v, last id = %v\n", rows, lastid)

	return nil
}

func (d *SqliteDatabase) GetVaultBlob(id string) ([]byte, error) {
	var blob []byte

	fmt.Printf("[db get vault] id: %v\n", id)

	if err := d.db.QueryRow(queryVaultStmt, id).Scan(&blob); err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("[db get vault] vault is empty\n")
			return nil, nil
		}

		return nil, err
	}

	return blob, nil
}

func (d *SqliteDatabase) SetVaultBlob(id string, blob []byte) error {
	_, err := d.db.Exec(insertOrUpdateVaultStmt, id, blob, blob)

	return err
}

func (d *SqliteDatabase) Close() error {
	return d.db.Close()
}


func (d *SqliteDatabase) bootstrapSchema() error {
	tx , err := d.db.Begin()

	if err != nil {
		return err
	}

	if _, err := tx.Exec(createVerifierTableStmt); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v\n", err, rbErr)
		}

		return err
	}

	if _, err := tx.Exec(createVerifierIndexStmt); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v\n", err, rbErr)
		}
		return err
	}

	if _, err := tx.Exec(createVaultTableStmt); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v\n", err, rbErr)
		}
		return err
	}

	if _, err := tx.Exec(createVaultIndexStmt); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v\n", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
