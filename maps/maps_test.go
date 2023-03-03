package maps_test

import (
	"fmt"
	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/maps"
	"testing"
)

func TestEachMap(t *testing.T) {
	m := map[string]interface{}{}
	m["2022"] = "2022"
	m["2021"] = "2021"
	m["2023"] = "2023"

	rsp := make(map[string]interface{}, 0)
	maps.SortMapKey(m, func(key string, value interface{}) {
		fmt.Println(key, value)
		rsp[key] = value
	}, 1)

	dumps.Dump(rsp)
}
