package gocheck

var timeoutMilliseconds int = 5000

type barrierResp struct {
	Err    error
	Resp   string
	Status int
}

func barrier(endpoints ...string) ([]int, []error) {

	requestNumber := len(endpoints)
	var status []int
	var responseError []error

	in := make(chan barrierResp, requestNumber)
	defer close(in)

	responses := make([]barrierResp, requestNumber)

	for _, endpoint := range endpoints {
		go statusCode(in, endpoint)
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

func StatusOK(endpoints ...string) string {
	responses, errors := barrier(endpoints...)

	for _, err := range errors {
		if err != nil {
			return "ko"
		}
	}

	for _, resp := range responses {
		if resp != 200 {
			return "ko"
		}
	}

	return "ok"

}
