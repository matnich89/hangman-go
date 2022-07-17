package util_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"hangman/util"
)

var _ = Describe("Word", func() {

	It("should return a random word", func() {
		Expect(util.GenerateWord()).To(Not(BeNil()))
	})

})
