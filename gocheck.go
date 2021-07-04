package gocheck

func StatusOK(endpoints ...string) string {
	responses, errors := BarrierStatusCode(endpoints...)

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
