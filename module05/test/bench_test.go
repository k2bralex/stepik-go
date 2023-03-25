package test

import (
	"sync"
	"testing"
)

type Counter struct {
	A int
	B int
}

func Increment(c *Counter) {
	c.A++
	c.B++
}

var pool = sync.Pool{New: func() interface{} { return new(Counter) }}

func BenchmarkWithoutPool(b *testing.B) {
	var c *Counter
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			c = &Counter{A: 1, B: 1}
			b.StopTimer()
			Increment(c)
			b.StartTimer()
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var c *Counter
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			c = pool.Get().(*Counter)
			c.A = 1
			c.B = 1
			b.StopTimer()
			Increment(c)
			b.StartTimer()
			pool.Put(c)
		}
	}
}
