package primes

import (
	"testing"
)

var (
	from = 1000
	to   = 4000
)

func BenchmarkUint16IsPrime(b *testing.B) {
	for n := uint16(from); n < uint16(to); n++ {
		Uint16IsPrime(n)
	}
}
func BenchmarkSieve(b *testing.B) {
	for m := from; m < to; m++ {
		Sieve(m)
	}
}
