package fetch

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	goUrl "net/url"
	"time"
)

type GetConfig struct {
	Headers map[string]string
	Params  map[string]string
}

type PostConfig struct {
	Headers map[string]string
	data    map[string]string
}

func Get(url string, options ...GetConfig) []byte {
	uri, err := goUrl.Parse(url)
	if err != nil {
		panic(err)
	}

	v := goUrl.Values{}

	config := options[0]
	if len(config.Params) != 0 {
		for key, value := range config.Params {
			v.Add(key, value)
		}
	}
	uri.RawQuery = v.Encode()

	request, err := http.Get(uri.String())
	if err != nil {
		panic(err)
	}

	if len(config.Headers) != 0 {
		for key, value := range config.Headers {
			request.Header.Set(key, value)
		}
	}

	defer request.Body.Close()

	data, _ := ioutil.ReadAll(request.Body)

	return data
}

func Post(url string, data interface{}, options ...PostConfig) []byte {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}

	config := options[0]
	if len(config.Headers) != 0 {
		for key, value := range config.Headers {
			req.Header.Set(key, value)
		}
	}

	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	response, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	result, _ := ioutil.ReadAll(response.Body)

	return result
}
