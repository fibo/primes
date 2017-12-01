package primes

func Uint16IsPrime(n uint16) bool {
	isPrime := map[uint16]bool{
		2:   true,
		3:   true,
		5:   true,
		7:   true,
		11:  true,
		13:  true,
		17:  true,
		19:  true,
		23:  true,
		29:  true,
		31:  true,
		37:  true,
		41:  true,
		43:  true,
		47:  true,
		53:  true,
		59:  true,
		61:  true,
		67:  true,
		71:  true,
		73:  true,
		79:  true,
		83:  true,
		89:  true,
		97:  true,
		101: true,
		103: true,
		107: true,
		109: true,
		113: true,
		127: true,
		131: true,
		137: true,
		139: true,
		149: true,
		151: true,
		157: true,
		163: true,
		167: true,
		173: true,
		179: true,
		181: true,
		191: true,
		193: true,
		197: true,
		199: true,
		211: true,
		223: true,
		227: true,
		229: true,
		233: true,
		239: true,
		241: true,
		251: true,
	}

	return isPrime[n]
}