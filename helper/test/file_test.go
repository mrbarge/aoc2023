package test

import (
	"github.com/mrbarge/aoc2023/helper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("File helpers", func() {

	Context("ReadLines", func() {
		It("Behaves correctly", func() {
			in := `abcd
efgh
ijkl`
			ir := strings.NewReader(in)
			out := []string{"abcd", "efgh", "ijkl"}
			res, err := helper.ReadLines(ir, false)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(out))
		})
		It("Ignores empty lines", func() {
			in := `abcd

efgh`
			ir := strings.NewReader(in)
			out := []string{"abcd", "efgh"}
			res, err := helper.ReadLines(ir, true)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(out))
		})

	})

	Context("ReadLinesAsInt", func() {
		It("Behaves correctly", func() {
			in := `123
-456
0
3322351`
			ir := strings.NewReader(in)
			out := []int{123, -456, 0, 3322351}
			res, err := helper.ReadLinesAsInt(ir)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(out))
		})
		It("Ignores empty lines", func() {
			in := `123

0`
			ir := strings.NewReader(in)
			out := []int{123, 0}
			res, err := helper.ReadLinesAsInt(ir)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(out))
		})
	})
	Context("ReadLinesAsIntArray", func() {
		It("Behaves correctly", func() {
			in := `123
456
7892
`
			ir := strings.NewReader(in)
			out := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9, 2}}
			res, err := helper.ReadLinesAsIntArray(ir)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(out))
		})
	})
})
