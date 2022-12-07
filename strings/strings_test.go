package strings_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/strings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimStr(t *testing.T) {
	s := strings.TrimStr(" ssss \t")
	dumps.Dump(s)
	assert.Equal(t, "ssss", s)
}

func TestAddPrefix(t *testing.T) {
	s := strings.AddPrefix(2, 5, "0")
	dumps.Dump(s)
	assert.Equal(t, "00002", s)
}

func TestCamelToCase(t *testing.T) {
	a := strings.CamelToCase("TestAll")
	dumps.Dump(a)
	assert.Equal(t, "test_all",a)
}

func TestCaseToCamel(t *testing.T) {
	a := strings.CaseToCamel("test_all")
	dumps.Dump(a)
	assert.Equal(t, "TestAll", a)
}
