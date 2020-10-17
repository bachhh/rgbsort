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
//			hilbert: using a hilbert walk in the colour space to linearize the colour values.
//			pointwise :
//			peano :
//			zcurve :
//
//
//
//

func main() {
	fmt.Println("vim-go")
}
