package uuids_test

import (
	"fmt"
	"testing"

	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/uuids"
	"github.com/stretchr/testify/assert"
)

func TestUuid(t *testing.T) {
	dumps.Dump(uuids.Uuid())
}

func TestGetRandowUUID(t *testing.T) {
	for i := 0; i < 1000; i++ {
		uuid := uuids.GetRandowUUID()
		fmt.Println(uuid)
		assert.NotEmpty(t, uuid)
	}
}

func TestGenerateUuid(t *testing.T) {
	for i := 0; i < 1000; i++ {
		uuid := uuids.GenerateUuid()
		fmt.Println(uuid)
		assert.NotEmpty(t, uuid)
		assert.Equal(t, 19, len(uuid))
	}
}

func TestGenerateUniqueNonceStr(t *testing.T) {
	str, err := uuids.GenerateUniqueNonceStr(16)
	if err != nil {
		assert.Error(t, err)
	}

	dumps.Dump(str)
}

func TestGenerateNonceStr(t *testing.T) {
	str, err := uuids.GenerateNonceStr(16)
	if err != nil {
		assert.Error(t, err)
	}

	dumps.Dump(str)
}
