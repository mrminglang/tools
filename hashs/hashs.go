package hashs

import (
	"crypto/sha1"
	"encoding/hex"
)

// 对字符串进行SHA1签名
func GenerateSHA1Signature(input string) string {
	// Convert the input string to bytes
	data := []byte(input)

	// Create a new SHA1 hash
	hash := sha1.New()

	// Write the data to the hash
	hash.Write(data)

	// Get the hashed result
	signature := hash.Sum(nil)

	// Convert the hashed result to a hexadecimal string
	signatureHex := hex.EncodeToString(signature)

	return signatureHex
}
