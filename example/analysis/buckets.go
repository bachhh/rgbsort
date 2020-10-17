package analysis

import (
	"sort"
	"sync/atomic"
)

type Buckets struct {
	b      []float64
	counts []uint64
}

// stolen from  github.com/prometheus/client_golang/blob/master/prometheus/histogram.go
func LinearBuckets(start, width float64, count int) (bkts *Buckets) {
	if count < 1 {
		panic("LinearBuckets needs a positive count")
	}
	buckets := make([]float64, count)
	for i := range buckets {
		buckets[i] = start
		start += width
	}
	return &Buckets{b: buckets, counts: make([]uint64, len(buckets))}
}

func ExponentialBuckets(start, factor float64, count int) *Buckets {
	if count < 1 {
		panic("ExponentialBuckets needs a positive count")
	}
	if start <= 0 {
		panic("ExponentialBuckets needs a positive start value")
	}
	if factor <= 1 {
		panic("ExponentialBuckets needs a factor greater than 1")
	}
	buckets := make([]float64, count)
	for i := range buckets {
		buckets[i] = start
		start *= factor
	}
	return &Buckets{b: buckets, counts: make([]uint64, len(buckets))}
}

func (b *Buckets) Observe(v float64) {
	atomic.AddUint64(&b.counts[b.find(v)], 1)
}

func (b *Buckets) find(v float64) int {
	return sort.SearchFloat64s(b.b, v)
}
