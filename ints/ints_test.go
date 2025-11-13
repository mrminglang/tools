package ints_test

import (
	"fmt"
	"testing"

	"github.com/mrminglang/tools/ints"
)

func TestInterfaceToInt32(t *testing.T) {
	// 测试不同类型转换
	values := []interface{}{
		int32(123),
		int(456),
		int64(789),
		float64(100.0),
		"1",
	}

	for _, v := range values {
		result, err := ints.InterfaceToInt32(v)
		if err != nil {
			fmt.Printf("Error converting %v: %v\n", v, err)
		} else {
			fmt.Printf("Converted %v to int32: %d\n", v, result)
		}
	}
}
