package util

import (
	"bytes"
	"com.azure.iot/iotcentral/iotcgo/config"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func GetContent(app string, url string) ([]byte, error) {
	appConfig, err := config.GetAppConfig(app)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", appConfig.ApiToken)

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func GetIndentedJSONContent(app string, url string) ([]byte, error) {
	body, err := GetContent(app, url)
	if err != nil {
		return nil, err
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "  ")
	if err != nil {
		return nil, err
	}

	return prettyJSON.Bytes(), nil
}

func PutContent(app string, url string, content []byte) ([]byte, error) {
	return httpCall(http.MethodPut, app, url, content)
}

func PostContent(app string, url string, content []byte) ([]byte, error) {
	return httpCall(http.MethodPost, app, url, content)
}

func httpCall(method string, app string, url string, content []byte) ([]byte, error) {
	appConfig, err := config.GetAppConfig(app)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", appConfig.ApiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil

}