package primes

import (
	"testing"
)

func TestUint16IsPrime(t *testing.T) {
	cases := []struct {
		num uint16
	}{
		{11171},
		{17},
	}

	for _, prime := range cases {
		ok := Uint16IsPrime(prime.num)

		if !ok {
			t.Errorf("Uint16IsPrime(%v) is false", prime.num)
		}
	}
}
