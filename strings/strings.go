package strings

import (
	"bytes"
	"fmt"
	"github.com/mrminglang/tools/checks/regex"
	"regexp"
	"strconv"
	"strings"
	"time"
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

// WhereInString where in String
func WhereInString(slices []string) string {
	str := ""
	if count := len(slices); count > 0 {
		for i := 0; i < count; i++ {
			str += "?,"
		}
		str = strings.Trim(str, ",")
	}

	return str
}

// WhereInStringValue where in String Value
func WhereInStringValue(args []interface{}, slices []string) []interface{} {
	if len(slices) <= 0 {
		return args
	}

	for _, row := range slices {
		args = append(args, row)
	}

	return args
}

// WhereInInt32 where in Int32
func WhereInInt32(slices []int32) string {
	str := ""
	if count := len(slices); count > 0 {
		for i := 0; i < count; i++ {
			str += "?,"
		}
		str = strings.Trim(str, ",")
	}

	return str
}

// WhereInInt32Value where in Int32 Value
func WhereInInt32Value(args []interface{}, slices []int32) []interface{} {
	if len(slices) <= 0 {
		return args
	}

	for _, row := range slices {
		args = append(args, row)
	}

	return args
}

// RebindSql 占位符替换
func RebindSql(dialect string, sqlstr *string) *string {
	switch dialect {
	case "mysql", "sqlite", "dm", "gbase", "clickhouse", "db2":
		return sqlstr
	}
	strs := strings.Split(*sqlstr, "?")
	if len(strs) < 1 {
		return sqlstr
	}
	var sqlBuilder strings.Builder
	sqlBuilder.WriteString(strs[0])
	for i := 1; i < len(strs); i++ {
		switch dialect {
		case "postgres", "postgresql", "kingbase": //postgres,postgresql,kingbase
			sqlBuilder.WriteString("$")
			sqlBuilder.WriteString(strconv.Itoa(i))
		case "mssql": //mssql
			sqlBuilder.WriteString("@p")
			sqlBuilder.WriteString(strconv.Itoa(i))
		case "oracle", "shentong": //oracle,神通
			sqlBuilder.WriteString(":")
			sqlBuilder.WriteString(strconv.Itoa(i))
		default: //其他情况,还是使用 ? | In other cases, or use  ?
			sqlBuilder.WriteString("?")
		}
		sqlBuilder.WriteString(strs[i])
	}
	*sqlstr = sqlBuilder.String()
	return sqlstr
}

// StringToSliceInt32 字符串转换为int切片
func StringToSliceInt32(str string, split string) []int32 {
	if str == "" {
		return nil
	}
	a := strings.Split(str, split)
	b := make([]int32, 0)
	for _, row := range a {
		c, _ := strconv.Atoi(row)
		b = append(b, int32(c))
	}
	return b
}

// GetStringYear 获取字符串中包含的年份
func GetStringYear(str string, len int) string {
	years := make([]int, 0)
	now := time.Now().Year()
	for i := 0; i < len; i++ {
		years = append(years, now-i)
	}

	var rsp int
	for _, v := range years {
		if strings.Contains(str, strconv.Itoa(v)) {
			rsp = v
			break
		}
	}

	if rsp <= 0 {
		return ""
	}
	return strconv.Itoa(rsp)
}

// StrUrlReplace 字符串中URL替换成指定格式
func StrUrlReplace(content, label string) string {
	re := regexp.MustCompile(regex.RegularUrl)
	urls := make(map[string]bool)
	// 使用map来记录已经替换过的URL，避免重复替换
	replaced := re.ReplaceAllStringFunc(content, func(url string) string {
		if _, ok := urls[url]; !ok {
			urls[url] = true
			return fmt.Sprintf(label, url)
		}
		return url
	})

	return replaced
}
