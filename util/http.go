package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const(
	BASE_TEST_URL = ""
	BASE_WEB_URL = "https://pay.akspay.com/api/v1.0"
	BASE_ERP_URL = "http://m.designjapan.cn/erp2"
	BASE_PRIZE_WEB_URL = "https://prize.utran.net"
)

func HTTPGet(uri string) ([]byte, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

//HTTPGet get request.
func NewHTTPGet(uri string, m map[string]string)([]byte, error){
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Set("timestamp", m["timestamp"])
	req.Header.Set("app-key", m["app-key"])
	req.Header.Set("sign", m["sign"])

	if err != nil {
		panic(err)
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

func NewHTTPPut(uri string, data string)([]byte, error){
	body := bytes.NewBuffer([]byte(data))
	request, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", uri, resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

func NewHTTPPost(uri string, data string)([]byte, error){
	body := bytes.NewBuffer([]byte(data))
	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", uri, resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

//HTTPPost post request.
func HTTPPost(uri string, data string)([]byte, error){
	body := bytes.NewBuffer([]byte(data))
	response, err := http.Post(uri, "", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}

//PostJSON post json data request.
func PostJSON(uri string, obj interface{})([]byte, error){
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u003c"), []byte("<"))
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u003e"), []byte(">"))
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u0026"), []byte("&"))
	body := bytes.NewBuffer(jsonData)
	response, err := http.Post(uri, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

//PostJSONWithResponseContentType post josn data and return data type.
func PostJSONWithResponseContentType(uri string, obj interface{})([]byte, string, error){
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, "", err
	}
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u003c"), []byte("<"))
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u003e"), []byte(">"))
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u0026"), []byte("&"))

	body := bytes.NewBuffer(jsonData)
	response, err := http.Post(uri, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")
	return responseData, contentType, err
}