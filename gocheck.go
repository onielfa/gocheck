package gocheck

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var timeoutMilliseconds int = 5000

func createConnection(url string) http.Client {
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}
	return client
}

//Check Status Code
func statusCode(url string) (int, error) {

	client := createConnection(url)
	response, err := client.Get(url)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return 0, err
	}
	return response.StatusCode, nil
}

//Check String response
func urlBody(url string) (string, error) {

	client := createConnection(url)
	response, err := client.Get(url)

	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	return string(body), nil

}

func StatusOK(url string) string {
	status, _ := statusCode(url)
	if status == 200 {
		return "ok"
	}
	return "ko"
}
