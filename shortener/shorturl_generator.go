package shortener

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

func base58Encoded(byte []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(byte)
	if err != nil {
		log.Panicf("Failed to encode: %v", err)
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortLink(initialLink, userId string) string {
	urlHashBytes := sha256Of(initialLink + userId)
	generatedNum := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finialStr := base58Encoded([]byte(fmt.Sprintf("%d", generatedNum)))
	return finialStr[:8]
}
