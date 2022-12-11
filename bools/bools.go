package bools

import "strings"

// TrueConverToBit 布尔true转1
func TrueConverToBit(sql string) string {
	return strings.Replace(sql, "true", "1", -1)
}

// FalseConverToBit 布尔false转0
func FalseConverToBit(sql string) string {
	return strings.Replace(sql, "false", "0", -1)
}

// BoolConverToBit b布尔转bit
func BoolConverToBit(sql string) string {
	sql = strings.Replace(sql, "true", "1", -1)
	sql = strings.Replace(sql, "false", "0", -1)
	return sql
}
