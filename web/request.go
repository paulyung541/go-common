package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"strings"
	"time"
)

const (
	TimeOut = time.Duration(20) * time.Second
)

// 通用的http请求
// method post 或者 get
// url 完整的url
// queryArgs query string 参数
// bodyArgs 请求的参数，根据不同的请求头被转成 form 或者 json 形式
// headers request请求头
// result 返回值为json，为指针类型
func doHttpHeader(method, url string, queryArgs, bodyArgs map[string]interface{}, headers map[string]string, result interface{}) error {
	var (
		req *http.Request
		err error
	)

	method = strings.ToUpper(method)
	if method != "GET" && method != "POST" {
		return errors.New("method must be get or post")
	}

	req, err = http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	if method == "POST" {
		if req.Header.Get("Content-Type") == "" {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			if bodyArgs != nil {
				form := netUrl.Values{}
				for key, value := range bodyArgs {
					form.Add(key, fmt.Sprint(value))
				}
			}
		}

		if req.Header.Get("Content-Type") == "application/json" && bodyArgs != nil {
			data, err := json.Marshal(bodyArgs)
			if err != nil {
				return err
			}
			req.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		}
	}

	queryParams := req.URL.Query()
	if queryArgs != nil {
		for key, value := range queryArgs {
			queryParams.Add(key, fmt.Sprint(value))
		}
	}

	req.URL.RawQuery = queryParams.Encode()

	client := &http.Client{Timeout: TimeOut}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err = resp.Body.Close(); nil != err {
			panic(err)
		}
	}()

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, result)
	if nil != err {
		return err
	}
	return nil
}

// DoGet GET请求
// queryArgs query string 参数
// result 返回值为json，为指针类型
func DoGet(url string, queryArgs map[string]interface{}, result interface{}) error {
	return doHttpHeader("get", url, queryArgs, nil, nil, result)
}

// DoPostForm POST请求，参数为表单 k1=v1&k2=v2...
// url 完整的url
// queryArgs query string 参数
// bodyArgs 请求的参数，会被转成json
// result 返回值为json，为指针类型
func DoPostForm(url string, queryArgs, bodyArgs map[string]interface{}, result interface{}) error {
	return doHttpHeader("post", url, queryArgs, bodyArgs, nil, result)
}

// DoPostJson POST请求，参数类型为json
// url 完整的url
// bodyArgs 请求的参数，会被转成json
// result 返回值为json，为指针类型
func DoPostJson(url string, queryArgs, bodyArgs map[string]interface{}, result interface{}) error {
	return doHttpHeader("post", url, queryArgs, bodyArgs, map[string]string{
		"Content-Type": "application/json",
	}, result)
}
