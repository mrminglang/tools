package strings

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// TrimStr 去除字符中一些前后字符
func TrimStr(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\n", "")
	return s
}

// AddPrefix 补全指定长度
func AddPrefix(v int, c int, flag string) string {
	s := strconv.Itoa(v)
	l := len(s)
	if l >= c {
		return s
	}
	prefix := strings.Repeat(flag, c-l)
	return fmt.Sprintf("%s%s", prefix, s)
}

// CamelToCase 驼峰写法转为下划线写法
func CamelToCase(name string) string {
	buffer := bytes.Buffer{}
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// CaseToCamel 下划线写法转为驼峰写法
func CaseToCamel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

