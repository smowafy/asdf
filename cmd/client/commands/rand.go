package commands

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"math/big"
)

var RandCommand = &cobra.Command{
	Use:   "rand",
	Short: "Generate a random password",
	Long: `
	Generate a random password containing uppercase and lowercase characters, numbers and symbols.
	`,
	Args: cobra.NoArgs,
	Run:  RunRandCommand,
}

func RunRandCommand(cmd *cobra.Command, args []string) {
	genPassword, err := generatePassword(passwordLength)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", genPassword)
}

const (
	symbolAlphabet    = "~`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/"
	lowercaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
	uppercaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numericAlphabet   = "0123456789"
	fullAlphabet      = "~`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var ErrRandPasswordShort = errors.New("Length must be at least 8")
var ErrRandPasswordLong = errors.New("Length must be at most 50")

func generatePassword(length int) (string, error) {
	if length < 8 {
		return "", ErrRandPasswordShort
	}

	if length > 50 {
		return "", ErrRandPasswordLong
	}

	passwordBytes := make([]byte, length)

	var idxBig, posBig *big.Int
	var err error

	posMaxBig := big.NewInt(int64(length))

	// add one character per type
	for _, alphabet := range []string{symbolAlphabet, lowercaseAlphabet, uppercaseAlphabet, numericAlphabet} {
		idxMaxBig := big.NewInt(int64(len(alphabet)))

		idxBig, err = rand.Int(rand.Reader, idxMaxBig)

		if err != nil {
			return "", err
		}

		for {
			posBig, err = rand.Int(rand.Reader, posMaxBig)

			if err != nil {
				return "", err
			}

			if passwordBytes[posBig.Int64()] == 0 {
				passwordBytes[posBig.Int64()] = alphabet[idxBig.Int64()]
				break
			}
		}
	}

	// generate the rest from the full alphabet
	for i := 0; i < length-4; i++ {
		idxMaxBig := big.NewInt(int64(len(fullAlphabet)))

		idxBig, err = rand.Int(rand.Reader, idxMaxBig)

		if err != nil {
			return "", err
		}

		for {
			posBig, err = rand.Int(rand.Reader, posMaxBig)

			if err != nil {
				return "", err
			}

			if passwordBytes[posBig.Int64()] == 0 {
				passwordBytes[posBig.Int64()] = fullAlphabet[idxBig.Int64()]
				break
			}
		}
	}

	return string(passwordBytes), nil
}
