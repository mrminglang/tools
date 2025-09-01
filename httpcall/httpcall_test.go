package httpcall_test

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"
	"testing"

	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/hashs"
	"github.com/mrminglang/tools/httpcall"
	"github.com/stretchr/testify/assert"
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

	body, err := httpcall.Post(reqUrl, nil, bytes.NewBuffer(jsonPayload))
	if err != nil {
		assert.Error(t, err)
	}

	dumps.Dump(body)
}

func TestPost_Json_2(t *testing.T) {
	reqUrl := "https://60.173.222.37:8443/apush/gateway/api/auth"
	loginId := "neican-yp"
	secret := "97861cdaa7272d066347dddb6589c148"
	payload := map[string]interface{}{
		"loginId": loginId,
		"sign":    hashs.GetAuthData(loginId, secret),
	}
	jsonPayload, _ := json.Marshal(payload)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	// 跳过TLS证书验证以测试自签名证书
	httpcall.SkipVerifyTLS = true
	body, err := httpcall.Post(reqUrl, headers, bytes.NewBuffer(jsonPayload))
	if err != nil {
		assert.Error(t, err)
	}

	dumps.Dump(body)

	// 恢复默认设置
	httpcall.SkipVerifyTLS = false
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

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	body, err := httpcall.Post("http://127.0.0.1:9701/coder", headers, payload)
	if err != nil {
		assert.Error(t, err)
	}

	dumps.Dump(body)
}
