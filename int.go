package primes

func IntIsPrime(n int) bool {
	isPrime := map[int]bool{
		2: true, 3: true, 5: true, 7: true, 11: true, 13: true, 17: true, 19: true, 23: true, 29: true, 31: true, 37: true, 41: true, 43: true, 47: true, 53: true, 59: true, 61: true, 67: true, 71: true, 73: true, 79: true, 83: true, 89: true, 97: true, 101: true, 103: true, 107: true, 109: true, 113: true, 127: true,
	}

	return isPrime[n]
}
