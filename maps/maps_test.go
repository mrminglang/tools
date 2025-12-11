package maps_test

import (
	"fmt"
	"testing"

	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/maps"
	"github.com/stretchr/testify/assert"
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

func TestConvertToMapStringSlice(t *testing.T) {
	class := map[string][]string{
		"A": {"A1"},
		"B": {"B1", "B2"},
	}

	m := map[string]interface{}{
		"car":   "car1",
		"class": class,
	}

	dumps.Dump(m)
	if _, ok := m["class"].(interface{}); ok {
		dumps.Dump(1111)
		rsp, err := maps.ConvertToMapStringSlice(m["class"])
		if err != nil {
			assert.Error(t, err)
		}
		dumps.Dump(rsp)
	} else {
		dumps.Dump(2222)
	}
}

func TestConvertToMapSliceString(t *testing.T) {
	param := map[string]string{
		"title":   "car1",
		"content": "内容",
	}
	rsp, err := maps.ConvertToMapSliceString(param)
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(rsp)
}
