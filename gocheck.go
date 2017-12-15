package gocheck

import "net/http"
import "fmt"

func checkRequest(url string) bool {

	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return false
	}

	fmt.Printf("Body: %v", response.Body)

	defer response.Body.Close()

	return true
}
