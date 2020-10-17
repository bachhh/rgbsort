package model

// Metric a function  that defines the concept of "distance" between any 2 members of a set,
// and in this space, "distance" between 2 colours
// Metric function should satisfies all the metric function properties:
//		d(a,b) == 0 <-> a == b
//		d(a,b) == d(b,a)
//		d(x,y) <= d(y,z) + d(z,x)
type Metric func(a, b []float64) float64
