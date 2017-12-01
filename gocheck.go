package gocheck

import "net/http"
import "fmt"

func checkRequest() bool {

	reponse, err := http.Get("http://www.google.com/")

	fmt.Printf("Body: %v", reponse.Body)
	fmt.Printf("Error: %v", err)

	if err == nil {
		return true
	}

	return false
}
