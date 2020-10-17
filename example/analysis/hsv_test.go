package analysis

import (
	"sort"
	"testing"

	"github.com/bachhh/rgbsort/model"
)

func TestHSV(t *testing.T) {
	t.Run("TestHSVDistribution", func(t *testing.T) {
		array := make([][]float64, 4)
		for i := range array {
			array[i] = []float64{}
		}
		for r := uint32(0); r < 256; r++ {
			for g := uint32(0); g < 256; g++ {
				for b := uint32(0); b < 256; b++ {
					c := &model.Colour{Red: r, Green: g, Blue: b}
					c.Model = model.HSV(c)
					// append 3 dimensions into our tabular
					for i := range c.Model {
						array[i] = append(array[i], c.Model[i])
						// the 4th dimensions comprised of all values from the other 3
						array[3] = append(array[3], c.Model[i])
					}
				}
			}
		}
		for i := range array {
			sort.Float64s(array[i])
		}

	})
}
