package uuids_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/times"
	"github.com/mrminglang/tools/uuids"
	"github.com/stretchr/testify/assert"
)

func TestUuid(t *testing.T) {
	dumps.Dump(uuids.Uuid())
}

func TestGetRandomUUID(t *testing.T) {
	for i := 0; i < 1000; i++ {
		uuid := uuids.GetRandomUUID()
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

func TestGetSimpleRandId(t *testing.T) {
	dumps.Dump(uuids.GetSimpleRandId())
}

func TestGetSimpleRandIdWithTime(t *testing.T) {
	now, _ := time.Parse(times.TimeFormat, "2025-09-15 23:59:59")
	dumps.Dump(uuids.GetSimpleRandIdWithTime(now))
}

func TestGetTaskId(t *testing.T) {
	dumps.Dump(uuids.GetTaskId())
}

func TestGetTaskIdWithTime(t *testing.T) {
	now, _ := time.Parse(times.TimeFormat, "2025-09-16 23:59:59")
	dumps.Dump(uuids.GetTaskIdWithTime(now))
}
