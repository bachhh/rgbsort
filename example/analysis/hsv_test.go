package analysis

import (
	"testing"

	"github.com/bachhh/rgbsort"
)

func TestHSV(t *testing.T) {
	t.Run("TestHSVDistribution", func(t *testing.T) {
		array := []*Colour{}
		for r := 0; r < 256; r++ {
			for g := 0; g < 256; g++ {
				for b := 0; b < 256; b++ {
					array = append(array, &rgbsort.Colour{Red: r, Green: g, Blue: b})
				}
			}
		}
	})
}
