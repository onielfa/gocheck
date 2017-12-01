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
			Expect(checkRequest()).Should(BeEquivalentTo(true))
		})
	})
})
