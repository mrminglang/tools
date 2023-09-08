package hashs_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/hashs"
	"testing"
)

func TestGenerateSHA1Signature(t *testing.T) {
	dumps.Dump(hashs.GenerateSHA1Signature("111"))
}

func TestMD5Hash(t *testing.T) {
	dumps.Dump(hashs.MD5Hash("2256-110,2257-111"))
}
