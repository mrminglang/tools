package pages

import (
	"fmt"
	"strings"
)

// GetPages 分页标识替换
func GetPages(sql string, page, size int) string {
	sql = strings.ReplaceAll(sql, "$LIMIT_STRING", GetLimitString(page, size))
	sql = strings.ReplaceAll(sql, "$BETWEEN_STRING", GetBetweenString(page, size))
	sql = strings.ReplaceAll(sql, "$OFFSET_STRING", GetOffsetString(page, size))
	sql = strings.ReplaceAll(sql, "$SKIP_STRING", GetSkipString(page, size))
	return sql
}

// GetLimitString LIMIT分页处理
func GetLimitString(page, size int) string {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	return fmt.Sprintf("LIMIT %d, %d", (page-1)*size, size)
}

// GetBetweenString Oracle分页处理
func GetBetweenString(page, size int) string {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	startRow := (page-1)*size + 1
	endRow := page * size

	return fmt.Sprintf("BETWEEN %d AND %d", startRow, endRow)
}

// GetOffsetString mssql分页处理
func GetOffsetString(page, size int) string {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	return fmt.Sprintf("OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", (page-1)*size, size)
}

// GetSkipString mongo分页处理
func GetSkipString(page, size int) string {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	return fmt.Sprintf("skip=%d;limit=%d", (page-1)*size, size)
}
