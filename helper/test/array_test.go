package test

import (
	"github.com/mrbarge/aoc2023/helper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Array helpers", func() {

	Context("When converting an array of numbers", func() {
		It("Returns a correct representation", func() {
			in := []string{"12", "5", "-6", "0", "81293912"}
			out := []int{12, 5, -6, 0, 81293912}

			r, err := helper.StrArrayToInt(in)
			Expect(err).To(BeNil())
			Expect(r).To(Equal(out))
		})

		It("Returns an error if an element is not an integer", func() {
			in := []string{"12", "five"}
			_, err := helper.StrArrayToInt(in)
			Expect(err).NotTo(BeNil())

		})
		It("Handles an empty array", func() {
			in := []string{}
			r, err := helper.StrArrayToInt(in)
			Expect(err).To(BeNil())
			Expect(r).To(BeEmpty())
		})
	})

	Context("Checking for int in array", func() {
		It("Behaves correctly", func() {
			in := []int{1, 2, 3, 4, -5, 6, 7, 8}
			Expect(helper.ContainsInt(1, in)).To(BeTrue())
			Expect(helper.ContainsInt(8, in)).To(BeTrue())
			Expect(helper.ContainsInt(4, in)).To(BeTrue())
			Expect(helper.ContainsInt(-5, in)).To(BeTrue())
			Expect(helper.ContainsInt(9, in)).To(BeFalse())
		})
	})

	Context("Converting string CSV to ints", func() {
		It("Behaves correctly", func() {
			in := "1,2,3,4,3,2,1"
			Expect(helper.StrCsvToIntArray(in, ",")).To(Equal([]int{1, 2, 3, 4, 3, 2, 1}))
		})
	})

})
