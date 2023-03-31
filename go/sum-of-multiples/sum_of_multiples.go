package summultiples

func SumMultiples(limit int, divisors ...int) int {

	multiples := make(map[int]int)

	for _, d := range divisors {
		if d == 0 {
			continue
		}
		for i := 1; d*i < limit; i++ {
			multiples[d*i]++
		}
	}
	sumofMultiples := 0

	for k := range multiples {
		sumofMultiples += k
	}

	return sumofMultiples
}
