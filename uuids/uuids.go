package uuids

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"strings"
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
