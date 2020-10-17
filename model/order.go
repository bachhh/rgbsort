package model

import (
	"fmt"

	"github.com/jtejido/hilbert"
)

// Order defines a total order on the n-dimensional real coordinates,
//	Returns true if a <= b, false otherwise
//
// To check if A and B are equivalance under the order:
//		Order(a, b) && Order (b, a)
// Order function should satisfies all the total order properties:
//		a <= a
//		(a <= b && b <= a ) <-> a == b
//		(a <= b && b <= c ) <-> a <= c
//		for any pair of colours a, b, either a <= b or b <= a
//
type Order func(a, b []float64) bool

// PointWise define the point-wise total order on the coordinate sapce
// Given 2 tuples a, b of type (K1, K2, ..., Kn):
//		a <= b iff either
//				(Kia < Kib)
//			  or (Kia == Kib) && (K(i+1)a <= K(i+1)b)... until kn
// PointWise will panic if a, b are not comparable (e.g.: different dimension)
func PointWise(a, b []float64) bool {
	// 2 coordinate space must have same dimension
	if len(a) != len(b) {
		panic(fmt.Sprintf("coordinates must have same dimension:  dim(a) = %d , dim(b) = %d", len(a), len(b)))
	}
	for i := range a {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	return true // a == b
}

// TODO ( bach ):
// HilbertPath construct a walking path for David Hilbert
// A HilbertWalk is a total order comparision by transform n-dimensioncoordinate to 1 dimension
// using a Hilbert curve
// Paramters:
//		- dim : dimension of the hypercube
//		- size: size of hypercube ( power of 2 )
//		- order: order of
func HilbertPath(dim, size, order int) (HilbertWalk Order, err error) {
	garden := hilbert.New(5, 2)
	return
}

// TODO ( bach ):
// SalesperonsWalk sort colour by using a traveling-salesman walk,
// distance between colouring point is calculated using a metric function
func SalespersonWalk(metricFunc Metric) (walker Order, err error) {
	return
}
