package network

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/emekarr/coding-test-busha/logger"
	"go.uber.org/zap"
)

type NetworkController struct {
	BaseUrl    string
	HttpClinet *http.Client
}

func (network *NetworkController) InitialiseNetworkClient() {
	network.HttpClinet = &http.Client{}
}

func (network *NetworkController) Get(path string, headers *map[string]string, params *map[string]string) *NetworkResponse {
	if network.HttpClinet == nil {
		network.InitialiseNetworkClient()
	}
	req, err := http.NewRequest("GET", network.BaseUrl+path, nil)
	if err != nil {
		logger.Error(fmt.Errorf("get request to %s%s failed", network.BaseUrl, path), zap.Error(err))
		return &NetworkResponse{
			Body:       nil,
			StatusCode: 400,
			Error:      err,
		}
	}
	setHeaders(headers, req)
	setParams(params, req)
	res, err := network.HttpClinet.Do(req)
	if err != nil {
		logger.Error(fmt.Errorf("get request to %s%s failed", network.BaseUrl, path), zap.Error(err))
		return &NetworkResponse{
			Body:       nil,
			StatusCode: res.StatusCode,
			Error:      err,
		}
	}
	res_body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Error(fmt.Errorf("get request to %s%s failed", network.BaseUrl, path), zap.Error(err))
		return &NetworkResponse{
			Body:       nil,
			StatusCode: res.StatusCode,
			Error:      err,
		}
	}
	res_json := string(res_body)
	return &NetworkResponse{
		Body:       res_json,
		StatusCode: res.StatusCode,
		Error:      nil,
	}
}

func (network *NetworkController) Post(path string, headers *map[string]string, body *map[string]interface{}, params *map[string]string) *NetworkResponse {
	if network.HttpClinet == nil {
		network.InitialiseNetworkClient()
	}
	parsed_body, err := json.Marshal(body)
	if err != nil {
		logger.Error(errors.New("error converting body to JSON"), zap.Error(err))
		return &NetworkResponse{
			Body:       nil,
			StatusCode: 400,
			Error:      err,
		}
	}
	req, err := http.NewRequest("POST", network.BaseUrl+path, bytes.NewBuffer(parsed_body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		logger.Error(fmt.Errorf("post request to %s%s failed", network.BaseUrl, path), zap.Error(err))
		return &NetworkResponse{
			Body:       nil,
			StatusCode: 400,
			Error:      err,
		}
	}
	setHeaders(headers, req)
	setParams(params, req)
	defer req.Body.Close()
	res, err := network.HttpClinet.Do(req)
	if err != nil {
		logger.Error(fmt.Errorf("post request to %s%s failed", network.BaseUrl, path), zap.Error(err))
		return &NetworkResponse{
			Body:       nil,
			StatusCode: res.StatusCode,
			Error:      err,
		}
	}
	res_body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Error(fmt.Errorf("post request to %s%s failed", network.BaseUrl, path), zap.Error(err))
		return &NetworkResponse{
			Body:       nil,
			StatusCode: res.StatusCode,
			Error:      err,
		}
	}
	res_json := string(res_body)
	return &NetworkResponse{
		Body:       res_json,
		StatusCode: res.StatusCode,
		Error:      nil,
	}
}

func setHeaders(headers *map[string]string, req *http.Request) {
	if headers == nil {
		return
	}
	for k := range *headers {
		req.Header.Add(k, (*headers)[k])
	}
}

func setParams(params *map[string]string, req *http.Request) {
	if params == nil {
		return
	}
	q := req.URL.Query()
	for k := range *params {
		q.Add(k, (*params)[k])
	}
	req.URL.RawQuery = q.Encode()
}
