package strings_test

import (
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/mrminglang/tools/dumps"
	strings2 "github.com/mrminglang/tools/strings"
	"github.com/stretchr/testify/assert"
)

func TestTrimStr(t *testing.T) {
	s := strings2.TrimStr(" ssss \t")
	dumps.Dump(s)
}

func TestTrimSignStr(t *testing.T) {
	pri, err := ioutil.ReadFile("./user_private.key")
	if err != nil {
		assert.Error(t, err)
	}
	pub, err := ioutil.ReadFile("./user_public.key")
	if err != nil {
		assert.Error(t, err)
	}

	userPrivate := strings2.TrimSignStr(string(pri))
	dumps.Dump(userPrivate)

	userPublic := strings2.TrimSignStr(string(pub))
	dumps.Dump(userPublic)

}

func TestAddPrefix(t *testing.T) {
	s := strings2.AddPrefix(2, 5, "0")
	dumps.Dump(s)
	assert.Equal(t, "00002", s)
}

func TestCamelToCase(t *testing.T) {
	a := strings2.CamelToCase("TestAll")
	dumps.Dump(a)
	assert.Equal(t, "test_all", a)
}

func TestCaseToCamel(t *testing.T) {
	a := strings2.CaseToCamel("test_all")
	dumps.Dump(a)
	assert.Equal(t, "TestAll", a)
}

func TestGetStringYear(t *testing.T) {
	title := "贵州茅台2023年第三季度报告"
	dumps.Dump(strings2.GetStringYear(title, 5))
}

func TestStrUrlReplace(t *testing.T) {
	content := "链接URL一:https://www.baid.com?a=a 链接URL三:https://www.baid.com?b=b 链接URL二:https://www.baid.com"
	label := "#link{%s}link#"
	content = strings2.StrUrlReplace(content, label)
	dumps.Dump(content)
}

func TestConvertToString(t *testing.T) {
	strs := []string{"1030032215", "1030027809", "1030020860", "1030031884"}
	result := strings2.ConvertToString(strs)
	dumps.Dump(result)
}

func TestIsBoolType(t *testing.T) {
	ok := strings2.IsBoolType("false")
	dumps.Dump(ok)
}

// IsTimestamp 判断是否为时间戳(秒级或毫秒级)
func IsTimestamp(timestamp int64) bool {
	length := len(strconv.FormatInt(timestamp, 10))

	return length == 10 || length == 13
}

// ConvertTimestamp 秒级时间戳转毫秒级时间戳
func ConvertTimestamp(timestamp int64) int64 {
	length := len(strconv.FormatInt(timestamp, 10))
	if length == 10 {
		return timestamp * 1000 // 转换为毫秒级时间戳
	}

	return timestamp // 否则返回原值
}

// ConvertToSeconds 毫秒级时间戳转秒级时间戳
func ConvertToSeconds(timestamp int64) int64 {
	length := len(strconv.FormatInt(timestamp, 10))
	if length == 13 {
		return timestamp / 1000 // 转换为秒级时间戳
	}

	return timestamp // 否则返回原值
}

func TestDisplay(t *testing.T) {
	req := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "mm",
		Age:  18,
	}

	dumps.Dump(strings2.Display(req))

}
