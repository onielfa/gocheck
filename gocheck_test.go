package gocheck

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGocheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go check Suite")
}

var _ = Describe("Go check Suite", func() {
	Context("Google get correct output", func() {

		It("Returns 200", func() {
			endpoints := []string{"http://www.google.com", "http://www.flywire.com"}
			result, _ := BarrierStatusCode(endpoints...)
			Expect(result[0]).To(Equal(200))
			Expect(result[1]).To(Equal(200))
		})

		It("one contains 403", func() {
			endpoints := []string{"http://www.google.com", "https://www.flywire.com/fsdfdsf"}
			result, _ := BarrierStatusCode(endpoints...)
			Expect(result[0]).To(Equal(200))
			Expect(result[1]).To(Equal(403))
		})

		It("fails when there is a timeout", func() {
			timeoutMilliseconds = 5
			endpoints := []string{"http://www.google.com"}
			_, responseError := BarrierStatusCode(endpoints...)
			Expect(responseError[0].Error()).To(ContainSubstring("Timeout"))
			timeoutMilliseconds = 5000
		})

		It("returns ok", func() {
			status := StatusOK("https://www.google.com")
			Expect(status).To(Equal("ok"))
		})

		It("returns ko", func() {
			status := StatusOK("https://www.google.com", "https://www.flywire.com/ferde")
			Expect(status).To(Equal("ko"))
		})

	})
})
