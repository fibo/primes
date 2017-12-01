package primes

import (
	"testing"
)

func TestUint8IsPrime(t *testing.T) {
	cases := []struct {
		num uint8
	}{
		{11},
		{17},
	}

	for _, prime := range cases {
		ok := Uint8IsPrime(prime.num)

		if !ok {
			t.Errorf("Uint8IsPrime(%v) is false", prime.num)
		}
	}
}
