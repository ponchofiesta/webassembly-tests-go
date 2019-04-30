package benchmarks

func Prime(max uint64) []uint64 {

	var primes []uint64

	if max > 2 {
		primes = append(primes, 2)
	}
	if max > 3 {
		primes = append(primes, 3)
	}

	sieve := make([]bool, max)
	for i := range sieve {
		sieve[i] = false
	}

	for x := uint64(1); x*x < max; x++ {
		for y := uint64(1); y*y < max; y++ {

			// Main part of Sieve of Atkin
			n := (4 * x * x) + (y * y)
			if n <= max && (n%12 == 1 || n%12 == 5) {
				sieve[n] = true
			}

			n = (3 * x * x) + (y * y)
			if n <= max && n%12 == 7 {
				sieve[n] = true
			}

			n = (3 * x * x) - (y * y)
			if x > y && n <= max && n%12 == 11 {
				sieve[n] = true
			}
		}
	}

	// Mark all multiples of squares as non-prime
	for r := uint64(5); r*r < max; r++ {
		if sieve[r] {
			for i := r * r; i < max; i += r * r {
				sieve[i] = false
			}
		}
	}

	// Print primes using sieve[]
	for a := uint64(5); a < max; a++ {
		if sieve[a] {
			primes = append(primes, a)
		}
	}

	return primes
}
