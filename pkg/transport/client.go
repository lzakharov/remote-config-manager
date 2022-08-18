package transport

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

func ListKeys() ([]string, error) {
	client := resty.New().SetTransport(http.DefaultTransport)

	resp, err := client.R().
		SetResult(&ListKeysResp{}).
		Get("http://127.0.0.1:8081/api/list")
	if err != nil {
		return nil, err
	}

	res := resp.Result().(*ListKeysResp)

	return res.Keys, nil
}

func Get(key string) (string, error) {
	client := resty.New().SetTransport(http.DefaultTransport)

	resp, err := client.R().
		SetQueryParam("key", key).
		Get("http://127.0.0.1:8081/api/get")

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func Put(key, value string) error {
	client := resty.New().SetTransport(http.DefaultTransport)

	_, err := client.R().
		SetQueryParam("key", key).
		SetBody(value).
		Post("http://127.0.0.1:8081/api/put")
	if err != nil {
		return err
	}

	return nil
}
