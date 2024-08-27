package cryptos_test

import (
	"fmt"
	"github.com/mrminglang/tools/cryptos"
	"github.com/mrminglang/tools/dumps"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {
	//timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	timestamp := 1724648633335
	plainText := fmt.Sprintf("%s|%d", "csyly04", timestamp)
	key := "pstock@%" // DES密钥必须为8字节
	iv := "upchina1"  // CBC模式的IV必须为8字节

	encrypted, err := cryptos.Encrypt(plainText, key, iv)
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(encrypted)
}

func TestDecrypt(t *testing.T) {
	encryptText := "WWuvkQFBTq/mZK80j8ahCJrFZZ0YGj14"
	key := "pstock@%" // DES密钥必须为8字节
	iv := "upchina1"  // CBC模式的IV必须为8字节

	decrypted, err := cryptos.Decrypt(encryptText, key, iv)
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(decrypted)
}
