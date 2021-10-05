package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiClient struct {
	baseUrl string
	port    string
}

func NewApiClient(baseUrl string, port string) ApiClient {
	return ApiClient{
		baseUrl: baseUrl,
		port:    port,
	}
}

func (a *ApiClient) Greet(name string) {
	url := fmt.Sprintf("http://%s:%s/%s", a.baseUrl, a.port, name)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(body))
		}
	}
}
