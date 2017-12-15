package gocheck

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoCheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gocheck")
}

var _ = Describe("Gocheck", func() {
	Context("Http connection", func() {
		It("Can connect", func() {
			Expect(checkRequest("https://www.google.com")).Should(BeEquivalentTo(true))
		})
		It("Can't connect", func() {
			Expect(checkRequest("https://google.com.cccc")).Should(BeEquivalentTo(false))
		})
	})
})
