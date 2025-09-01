package hashs

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"
)

// GenerateSHA1Signature 对字符串进行SHA1签名
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

// MD5Hash 对字符串MD5哈希加密
func MD5Hash(str string) string {
	hash := md5.Sum([]byte(str))
	md5Str := hex.EncodeToString(hash[:])
	return md5Str
}

// GetAuthData 生成认证数据签名
func GetAuthData(loginId, secret string) string {
	data := fmt.Sprintf("loginId%s%s", loginId, secret)
	return strings.ToUpper(Sha1Hex(data))
}

// Sha1Hex 计算字符串的 SHA1 哈希值并返回十六进制表示
func Sha1Hex(s string) string {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
