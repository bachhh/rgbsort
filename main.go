package main

import "fmt"

// RGBSort
// RGBSort will order colour space into the linear order ( chain ), using different orders
// Inspired by https://www.alanzucconi.com/2015/09/30/colour-sorting/

// As the colour space is multi dimensional, there is no 1 universally correct way
// to order. Each one will will have appeals and drawbacks depending on
// your needs, usages, and intuition.

// Flags
//		--format: specify format of hexed red-green-blue colours, using tokens. Default is #2R2G2B2A. ( hex-encoded rgbA )
//			tokens value:
//
// 		 	- 2R: red colour value in HEX
// 		 	- 2G: green colour value in HEX
// 		 	- 2B: blue colour value in HEX
// 		 	- 2A: alpha value in HEX
// 		 	- 3R: red colour value in base 256
// 		 	- 3G: green colour value in base 256
// 		 	- 3B: blue colour value in base 256
// 		 	- 3A: alpha value in base 256
//	[*] for every entries, missing RGB values are assumed zero. But for alpha it is assumed maximum value
//
// 		--order: specify the order for colours, supported values:
//
//			hilbert-curve : using a hilbert walk in the colour space to linearize the colour values.
//			luminosity :
//
//
//
//

func main() {
	fmt.Println("vim-go")
}

type Colour struct {
	Red      uint32
	Green    uint32
	Blue     uint32
	Alpha    uint32
	RawInput string
	// Space    []float64 // space defines the colour model
}

// Order defines a total order on the n-dimensional real coordinates,
//	Returns true if a <= b, false otherwise
//
// To check if A and B are equivalance under the order:
//		Order(a, b) && Order (b, a)
type Order func(a, b []float64) bool

// ColourSpace colour space defines a model function that maps Colours object
// onto n-dimension real coordinates ( usually [0, 1.0] )
type ColourSpace func(c *Colour) []float64
