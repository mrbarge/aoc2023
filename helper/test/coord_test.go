package test

import (
	"github.com/mrbarge/aoc2023/helper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Coord helpers", func() {

	Context("When getting neighbours", func() {
		It("behaves successfully", func() {
			in := helper.Coord{X: 0, Y: 0}
			out := []helper.Coord{{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1}, {X: -1, Y: 0}, {X: 1, Y: 0}, {X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1}}
			ret := in.GetNeighbours(true)
			for _, v := range out {
				Expect(ret).To(ContainElement(v))
			}
			Expect(len(ret)).To(Equal(len(out)))
		})
	})

	Context("When getting non-negative neighbours", func() {
		It("behaves successfully", func() {
			in := helper.Coord{X: 0, Y: 0}
			out := []helper.Coord{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}}
			ret := in.GetNeighboursPos(true)
			Expect(ret).To(ContainElements(out))
			Expect(len(ret)).To(Equal(len(out)))

			in = helper.Coord{X: 2, Y: 2}
			out = []helper.Coord{{X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}, {X: 2, Y: 1}, {X: 2, Y: 3}, {X: 1, Y: 3}, {X: 2, Y: 3}, {X: 3, Y: 3}}
			ret = in.GetNeighboursPos(true)

			for _, v := range out {
				Expect(ret).To(ContainElement(v))
			}
			Expect(len(ret)).To(Equal(len(out)))
		})
	})

})
