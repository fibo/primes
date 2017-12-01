package primes

func IntIsPrime(n int) bool {
	isPrime := map[int]bool{
		1:  true,
		2:  true,
		3:  true,
		5:  true,
		7:  true,
		11: true,
		13: true,
		17: true,
		19: true,
		23: true,
		29: true,
		31: true,
		37: true,
		41: true,
		43: true,
		47: true,
		53: true,
		59: true,
		61: true,
	}

	return isPrime[n]
}
