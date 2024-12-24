package shortener

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// sha256Of computes the SHA256 checksum of the given input string.
// It uses the crypto/sha256 package to compute the checksum, and
// returns the result as a byte slice.
func sha256Of(input string) []byte {
	algo := sha256.New()
	// Write the input string to the SHA256 algorithm.
	algo.Write([]byte(input))
	// Return the computed checksum as a byte slice.
	return algo.Sum(nil)
}

// base58Encoded takes a byte slice as input and returns its Base58-encoded representation.
// The function panics if encoding fails.
func base58Encoded(input []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(input)
	if err != nil {
		// If encoding fails, panic and exit with an error code.
		log.Panicf("Failed to encode: %v", err)
		os.Exit(1)
	}
	return string(encoded)
}

// GenerateShortLink generates a short link using the initial link and user ID as input.
// It computes the SHA256 hash of the concatenated input, converts it to a big integer,
// encodes it in Base58, and returns the first 8 characters of the encoded string.
func GenerateShortLink(initialLink, userId string) string {
	// Compute the SHA256 checksum of the concatenated initialLink and userId.
	urlHashBytes := sha256Of(initialLink + userId)

	// Convert the hash bytes to a big integer and get its unsigned 64-bit representation.
	generatedNum := new(big.Int).SetBytes(urlHashBytes).Uint64()

	// Encode the number as a Base58 string.
	finialStr := base58Encoded([]byte(fmt.Sprintf("%d", generatedNum)))

	// Return the first 8 characters of the encoded string as the short link.
	return finialStr[:8]
}
