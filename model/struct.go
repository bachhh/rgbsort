package model

type Colour struct {
	Red      uint32
	Green    uint32
	Blue     uint32
	Alpha    uint32
	RawInput string
	Model    []float64 // cache colour model value
}
