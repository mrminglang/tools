package httpcall

import (
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// SkipVerifyTLS 设置是否跳过TLS证书验证
var SkipVerifyTLS = false

// Get 发起一个Get请求
func Get(url string) (string, error) {
	body := new(io.Reader)
	logrus.Infoln("Get url::", url)
	return do(http.MethodGet, url, nil, *body)
}

// Post 发起一个Post请求
func Post(url string, headers map[string]string, body io.Reader) (string, error) {
	logrus.Infoln(fmt.Sprintf("Post url:: %s headers::%v", url, headers))
	logrus.Infoln("Post body::", body)
	return do(http.MethodPost, url, headers, body)
}

// 远程请求
func do(method, url string, headers map[string]string, reqBody io.Reader) (string, error) {
	logrus.Infoln("do creating request start......")
	if headers == nil {
		headers = make(map[string]string)
	}
	if len(headers) == 0 {
		headers["Content-Type"] = "application/json"
	}

	// 准备请求正文
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		logrus.Infoln("do creating request error::", err.Error())
		return "", err
	}
	// 设置“内容类型”标头
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// 创建HTTP客户端，支持跳过TLS验证
	client := &http.Client{}
	if SkipVerifyTLS {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	logrus.Infoln("do req::", req)
	rsp, err := client.Do(req)
	if err != nil {
		logrus.Errorln("do making request error::", err.Error())
		return "", err
	}
	defer rsp.Body.Close()

	// 读取响应正文
	var rspBody string
	switch rsp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(rsp.Body)
		buf := make([]byte, 1024)
		for {
			n, err := reader.Read(buf)
			if err != nil && err != io.EOF {
				logrus.Infof("read Buffer Error : %v", err)
				return "", err
			}
			if n == 0 {
				break
			}

			rspBody += string(buf)
		}
	default:
		bodyByte, _ := ioutil.ReadAll(rsp.Body)
		rspBody = string(bodyByte)
	}

	logrus.Infoln("do creating request end body::", rspBody)
	return rspBody, nil
}
