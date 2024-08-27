package cryptos

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

// Encrypt 加密
func Encrypt(plainText, key, iv string) (string, error) {
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	var mode cipher.BlockMode
	mode = cipher.NewCBCEncrypter(block, []byte(iv))

	// 填充
	paddedText := pad([]byte(plainText), block.BlockSize())
	cipherText := make([]byte, len(paddedText))
	mode.CryptBlocks(cipherText, paddedText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// 填充
func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

// Decrypt 解密
func Decrypt(encryptText, key, iv string) (string, error) {
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	var mode cipher.BlockMode
	mode = cipher.NewCBCDecrypter(block, []byte(iv))

	cipherText, err := base64.StdEncoding.DecodeString(encryptText)
	if err != nil {
		return "", err
	}

	plainText := make([]byte, len(cipherText))
	mode.CryptBlocks(plainText, cipherText)

	// 去除填充
	unPaddedText, err := unPad(plainText)
	if err != nil {
		return "", err
	}

	return string(unPaddedText), nil
}

// 去除填充
func unPad(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return nil, fmt.Errorf("input is empty")
	}
	padding := int(src[length-1])
	if padding > length {
		return nil, fmt.Errorf("padding size is invalid")
	}
	return src[:length-padding], nil
}
