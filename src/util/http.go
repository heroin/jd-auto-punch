package util

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"time"
)

var (
	agent = []string{
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1554.0 Safari/537.36",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0)",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.71 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.72 Safari/537.36",
	}
)

func Connect() {
	response, err := Client().Post("http://******/login", "",
		bytes.NewBufferString("username=******&password=******&login_type=login"))
	if err != nil {
		ERROR("open connect is ERROR, %s", err)
	} else {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		if body != nil {
			INFO("open connect SUCCESS!")
		}
	}
}

func Client() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, time.Duration(3)*time.Second)
			},
		},
	}
	return client
}

func GetUrlInUserAgent(url string) (resp *http.Response, err error) {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", RandomUserAgent())
	return Client().Do(request)
}

func RandomUserAgent() string {
	return agent[rand.Intn(len(agent))]
}
