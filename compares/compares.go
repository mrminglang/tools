package compares

import "strings"

// GetCompares 获取比较符号
func GetCompares(str string) string {
	str = strings.ReplaceAll(str, "$eq", "=")   // 等于
	str = strings.ReplaceAll(str, "$ne", "!=")  // 不等于
	str = strings.ReplaceAll(str, "$gte", ">=") // 大于等于
	str = strings.ReplaceAll(str, "$gt", ">")   // 大于
	str = strings.ReplaceAll(str, "$lte", "<=") // 小于等于
	str = strings.ReplaceAll(str, "$lt", "<")   // 小于
	return str
}
