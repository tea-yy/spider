package util

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

//获取HTTP请求返回数据
func GetHttpResponse(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("NET ERROR:" + strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	utfBody, err := GbkToUtf8(body)
	if err != nil {
		return "", err
	}
	return string(utfBody), nil
}
