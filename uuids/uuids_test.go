package uuids_test

import (
	"fmt"
	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/uuids"
	"github.com/stretchr/testify/assert"
	"testing"
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
