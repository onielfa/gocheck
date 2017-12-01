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
		It("returns 200", func() {
			status, _ := statusCode("https://www.google.com")
			Expect(status).To(Equal(200))
		})

		It("fails when there is a timeout", func() {
			timeoutMilliseconds = 5
			_, err := statusCode("https://www.google.com")
			Expect(err.Error()).To(ContainSubstring("Timeout"))
			timeoutMilliseconds = 5000
		})

		It("Contains text", func() {
			Expect(urlBody("https://www.google.com")).To(ContainSubstring("<title>Google</title>"))
		})

	})
})
