package main

import "math"

func RGBa(c *Colour) []float64 {
	return []float64{float64(c.Red), float64(c.Green), float64(c.Blue), float64(c.Alpha)}
}

// Map colour to HSV space
func HSV(c *Colour) []float64 {
	return []float64{Hue(c), SatHV(c), Value(c)}
}

// HSLSort same as as HSV but use a different value: lightness
func HSL(c *Colour) []float64 {
	return []float64{Hue(c), SatHL(c), Luma(c)}
}

// YIQ colour model used in analog television networks
func YIQ(c *Colour) []float64 {
	y := 0.30*float64(c.Red) + 0.59*float64(c.Green) + 0.11*float64(c.Blue)
	i := 0.74*(float64(c.Red)-y) - 0.27*(float64(c.Blue)-y)
	q := 0.48*(float64(c.Red)-y) + 0.41*(float64(c.Blue)-y)
	return []float64{y, i, q}
}

// TODO: CMYK the subtractive cyan, magenta, yellow, black colour model
func CMYK(c *Colour) []float64 {
	return nil
}

// Hue calculation
func Hue(a *Colour) (h float64) {
	ma, mi := max(a.Red, max(a.Green, a.Blue)), min(a.Red, min(a.Green, a.Blue))
	if ma != mi {
		if ma == a.Red {
			h = float64(a.Green-a.Blue) / float64(ma-mi)
		}
		if ma == a.Green {
			h = float64(a.Blue-a.Red)/float64(ma-mi) + 2.0
		}
		if ma == a.Blue {
			h = float64(a.Red-a.Green)/float64(ma-mi) + 4.0
		}
		h = math.Mod(h/6.0, 1.0)
		return
	}
	return 0
}

func Value(a *Colour) float64 {
	return float64(max(a.Red, max(a.Green, a.Blue)))
}

// Luma / Lightness
func Luma(a *Colour) float64 {
	return float64(max(a.Red, max(a.Green, a.Blue))+min(a.Red, min(a.Green, a.Blue))) / 2
}

// NOTE:  Saturation is the scale factor dependent on the colour model,
// the calculations are different between HSL and HSV
func SatHL(a *Colour) float64 {
	ma, mi := max(a.Red, max(a.Green, a.Blue)), min(a.Red, min(a.Green, a.Blue))
	if ma == mi {
		return 0
	}

	l := Luma(a)
	if lte(l, 1/2) {
		return float64(ma-mi) / (2 - 2*l)
	}
	return float64(ma-mi) / (2 * l)
}

func SatHV(a *Colour) float64 {
	ma, mi := max(a.Red, max(a.Green, a.Blue)), min(a.Red, min(a.Green, a.Blue))
	if ma == 0 {
		return 0
	}
	return 1 - float64(mi)/float64(ma)
}
