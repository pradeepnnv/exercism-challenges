package sieve

func Sieve(limit int) []int {
	r := make([]int, limit)

	// initialize the range

	for i := 1; i <= limit; i++ {
		if i > 1 {
			r[i-1] = i
		}
	}

	// iterate through the range of primes
	for _, prime := range r {
		if prime < 2 || prime == 0 {
			continue
		}

		for i := 2; i*prime <= limit; i++ {

			// if it's in the range..set to zero
			r[i*prime-1] = 0

		}
	}

	var finalList []int
	//remove zeros
	for _, v := range r[1:] {
		if v != 0 {
			finalList = append(finalList, v)
		}
	}
	return finalList

}
