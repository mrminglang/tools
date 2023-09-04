package hashs_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/hashs"
	"testing"
)

func TestGenerateSHA1Signature(t *testing.T) {
	dumps.Dump(hashs.GenerateSHA1Signature("111"))
}
