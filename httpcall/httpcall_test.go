package httpcall_test

import (
	"bytes"
	"encoding/json"
	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/httpcall"
	"github.com/stretchr/testify/assert"
	"net/url"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	// 百度地图
	body, err := httpcall.Get("https://api.map.baidu.com/reverse_geocoding/v3/?location=39.92594,116.37304&output=json&ak=")
	if err != nil {
		assert.Error(t, err)
	}

	dumps.Dump(body)
}

// 请求格式为JSON
func TestPost_Json(t *testing.T) {
	reqUrl := "http://127.0.0.1:9701/json/cnews_attach_fetch/getF10FinanceReport"
	payload := map[string]interface{}{
		"req": map[string]interface{}{
			"stockCode": "835368",
			"market":    7,
		},
	}
	jsonPayload, _ := json.Marshal(payload)

	body, err := httpcall.Post(reqUrl, "", bytes.NewBuffer(jsonPayload))
	if err != nil {
		assert.Error(t, err)
	}

	dumps.Dump(body)
}

// 请求格式为FORM
func TestPost_Form(t *testing.T) {
	params := url.Values{}
	params.Set("lat", "39.513668060302734")
	params.Set("lng", "106.63082122802734")
	params.Set("fixed", "0")
	params.Set("lang", "cn")

	// Prepare the request body
	payload := strings.NewReader(params.Encode())

	contentType := "application/x-www-form-urlencoded"
	body, err := httpcall.Post("http://127.0.0.1:9701/coder", contentType, payload)
	if err != nil {
		assert.Error(t, err)
	}

	dumps.Dump(body)
}
