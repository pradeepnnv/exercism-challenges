package prime

import "fmt"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("prime cannot be determined")
	}

	primes := []int{2, 3}

	if n < len(primes) {
		return primes[n-1], nil
	}
	fmt.Println(primes)

	for n > len(primes) {
		//identifying the next prime
		np := primes[len(primes)-1] + 1
		fmt.Println("np:", np)
		for _, v := range primes {
			fmt.Println("v:", v)
			if n%v == 0 {
				break
			}
			if v > np {
				primes = append(primes, np)
			}
		}
		fmt.Println("primes :", primes)
		np++
	}

	return primes[n-1], nil
}
