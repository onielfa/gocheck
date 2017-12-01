package gocheck

import (
	"io/ioutil"
	"net/http"
	"time"
)

func createConnection(url string) http.Client {
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}
	return client
}

//Check Status Code
func statusCode(out chan<- barrierResp, url string) {

	response := barrierResp{}

	client := createConnection(url)
	res, err := client.Get(url)

	if err != nil {
		response.Err = err
		out <- response
		return
	}

	response.Status = res.StatusCode
	out <- response
}

//Check String response
func urlRequest(out chan<- barrierResp, url string) {

	response := barrierResp{}
	client := createConnection(url)

	resp, err := client.Get(url)
	if err != nil {
		response.Err = err
		out <- response
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Err = err
		out <- response
		return
	}

	response.Resp = string(byt)
	out <- response

}
