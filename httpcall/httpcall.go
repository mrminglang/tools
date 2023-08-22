package httpcall

import (
	"compress/gzip"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

//Get 发起一个Get请求
func Get(url string) (string, error) {
	body := new(io.Reader)
	logrus.Infoln("Get url::", url)
	return do(http.MethodGet, url, "", *body)
}

//Post 发起一个Post请求
func Post(url, contentType string, body io.Reader) (string, error) {
	logrus.Infoln(fmt.Sprintf("Post url:: %s contentType::%s", url, contentType))
	logrus.Infoln("Post body::", body)
	return do(http.MethodPost, url, contentType, body)
}

// 远程请求
func do(method, url, contentType string, reqBody io.Reader) (string, error) {
	if contentType == "" {
		contentType = "application/json"
	}
	logrus.Infoln("do creating request start......")
	// 准备请求正文
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		logrus.Infoln("do creating request error::", err.Error())
		return "", err
	}
	// 设置“内容类型”标头
	req.Header.Add("Content-Type", contentType)
	// Make the request
	client := &http.Client{}
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
