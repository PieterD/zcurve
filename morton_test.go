package zcurve

import (
	"math/rand"
	"testing"
)

func BenchmarkInterleave(b *testing.B) {
	b.StopTimer()
	var coords = make([][2]uint32, b.N)
	for n := 0; n < b.N; n++ {
		coords[n] = [2]uint32{uint32(rand.Int63()), uint32(rand.Int63())}
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		interleaved := interleave(coords[n][0], coords[n][1])
		if interleaved < uint64(coords[n][0])+uint64(coords[n][1]) {
			b.Fatalf("interleaved smaller than sum of parts")
		}
	}
}

func BenchmarkDeinterleave(b *testing.B) {
	b.StopTimer()
	var coords = make([][2]uint32, b.N)
	for n := 0; n < b.N; n++ {
		coords[n] = [2]uint32{uint32(rand.Int63()), uint32(rand.Int63())}
	}
	var interleaved = make([]uint64, b.N)
	for n := 0; n < b.N; n++ {
		interleaved[n] = interleave(coords[n][0], coords[n][1])
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		x, y := deinterleave(interleaved[n])
		if x != coords[n][0] || y != coords[n][1] {
			b.Fatalf("invalid deinterleave %016x: expected %d,%d but got %d,%d", interleaved[n], coords[n][0], coords[n][1], x, y)
		}
	}
}
