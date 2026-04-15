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
		dumps.Dump(m["class"])
		rsp, err := maps.ConvertToMapStringSlice(m["class"])
		if err != nil {
			assert.Error(t, err)
		}
		dumps.Dump(2222)
		dumps.Dump(rsp)
	} else {
		dumps.Dump(333)
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

	param1 := "{\"name\":\"abc\"}"
	rsp1, err1 := maps.ConvertToMapSliceString(param1)
	if err1 != nil {
		assert.Error(t, err1)
	}
	dumps.Dump(rsp1)

}

func TestConvertToMapStringInterface(t *testing.T) {
	data := "{\"13800010001\":{\"src_id\":\"10001\",\"content\":\"666500000201899230\"},\"13800010002\":{\"src_id\":\"10002\",\"content\":\" 00000202079416+李**666\"}}"
	rsp, err := maps.ConvertToMapStringInterface(data)
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(rsp)

	data1 := map[string]interface{}{
		"13800010001": map[string]interface{}{
			"src_id":  "10001",
			"content": "666500000201899230",
		},
		"13800010002": map[string]interface{}{
			"src_id":  "10002",
			"content": " 00000202079416+李**666",
		},
	}
	rsp1, err1 := maps.ConvertToMapStringInterface(data1)
	if err1 != nil {
		assert.Error(t, err1)
	}
	dumps.Dump(rsp1)

	content := "{\"sms\":{\"content\":\"${name}_${content}\"}}"
	rsp2, err2 := maps.ConvertToMapStringInterface(content)
	if err2 != nil {
		assert.Error(t, err2)
	}
	dumps.Dump(rsp2)

	content2 := map[string]interface{}{
		"sms": map[string]interface{}{
			"title":   "${name}_${title}",
			"content": "${name}_${content}",
		},
	}
	rsp3, err3 := maps.ConvertToMapStringInterface(content2)
	if err3 != nil {
		assert.Error(t, err3)
	}
	dumps.Dump(rsp3)
}

func TestConvertToMapStringMap(t *testing.T) {
	content := "{\"sms\":{\"content\":\"${name}_${content}\"}}"
	rsp, err := maps.ConvertToMapStringMap(content)
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(rsp)
	dumps.Dump(rsp["sms"]["content"])

	content1 := map[string]interface{}{
		"sms": map[string]interface{}{
			"title":   "${name}_${title}",
			"content": "${name}_${content}",
		},
	}
	rsp1, err1 := maps.ConvertToMapStringMap(content1)
	if err1 != nil {
		assert.Error(t, err1)
	}
	dumps.Dump(rsp1)
	dumps.Dump(rsp1["sms"]["content"])
}
