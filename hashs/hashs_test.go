package hashs_test

import (
	"testing"

	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/hashs"
)

func TestGenerateSHA1Signature(t *testing.T) {
	dumps.Dump(hashs.GenerateSHA1Signature("111"))
}

func TestMD5Hash(t *testing.T) {
	dumps.Dump(hashs.MD5Hash("2256-110,2257-111"))
}

func TestGetAuthData(t *testing.T) {
	loginId := "neican-yp"
	secret := "97861cdaa7272d066347dddb6589c148"
	dumps.Dump(hashs.GetAuthData(loginId, secret))
	loginId2 := "kuaixun-yp"
	secret2 := "5bb1848fec4d76f64be533e34ea8ab14"
	dumps.Dump(hashs.GetAuthData(loginId2, secret2))
}
