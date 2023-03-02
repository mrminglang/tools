package maps_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/maps"
	"testing"
)

func TestSortKeyRise(t *testing.T) {
	req := map[string]interface{}{
		"2022": "3332",
		"2023": "2222",
		"2021": "1111",
	}

	rsp := maps.SortKeyRise(req)
	dumps.Dump(rsp)
}
