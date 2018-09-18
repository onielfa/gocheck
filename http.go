package gocheck

import (
	"io/ioutil"
	"net/http"
	"time"
)

var timeoutMilliseconds int = 5000

type barrierResp struct {
	Err    error
	Resp   string
	Status int
}

func BarrierStatusCode(endpoints ...string) ([]int, []error) {

	requestNumber := len(endpoints)
	var status []int
	var responseError []error

	in := make(chan barrierResp, requestNumber)
	defer close(in)

	responses := make([]barrierResp, requestNumber)

	for _, endpoint := range endpoints {
		go StatusCode(in, endpoint)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		responseError = append(responseError, resp.Err)
		if resp.Err != nil {
			hasError = true
		}
		responses[i] = resp
	}
	if !hasError {
		for _, resp := range responses {
			status = append(status, resp.Status)
		}
	}

	return status, responseError

}

func createConnection(url string) http.Client {
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}
	return client
}

//Check Status Code
func StatusCode(out chan<- barrierResp, url string) {

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
