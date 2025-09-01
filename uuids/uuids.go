package uuids

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

// GenerateUuid 雪花算法生成19位唯一id
func GenerateUuid() string {
	// Create a new Node with a Node number of 1
	node, _ := snowflake.NewNode(int64(1))

	// Generate a snowflake ID.
	id := node.Generate().String()
	return id
}

// Uuid 返回有连接符的UUID
func Uuid() string {
	return uuid.New().String()
}

// GetRandowUUID 返回无连接符的UUID
func GetRandowUUID() string {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	str := newUUID.String()
	strArray := strings.Split(str, "-")
	str = strings.Join(strArray, "")
	return str
}

// GenerateUniqueNonceStr 生成唯一性随机字符串
func GenerateUniqueNonceStr(length int) (string, error) {
	nonceStr, err := GenerateNonceStr(length)
	if err != nil {
		return "", err
	}
	// 添加时间戳作为唯一标识符
	timestamp := time.Now().UnixNano()
	nonceStr = fmt.Sprintf("%s%d", nonceStr, timestamp)
	return nonceStr, nil
}

// GenerateNonceStr 生成随机字符串
func GenerateNonceStr(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Use base64 encoding to generate the random bytes
	nonceStr := base64.URLEncoding.EncodeToString(randomBytes)

	// Trim the nonceStr to the specified length
	if len(nonceStr) > length {
		nonceStr = nonceStr[:length]
	}

	return nonceStr, nil
}
