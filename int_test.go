package primes

import (
	"testing"
)

func TestIntIsPrime(t *testing.T) {
	cases := []struct {
		num int
	}{
		{11},
		{17},
	}

	for _, prime := range cases {
		ok := IntIsPrime(prime.num)

		if !ok {
			t.Errorf("IntIsPrime(%v) is false", ok)
		}
	}
}
