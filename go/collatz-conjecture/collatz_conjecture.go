package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	c := 0
	if n <= 0 {
		return 0, fmt.Errorf("error")
	}

	for {
		if n == 1 {
			break
		} else {
			if n%2 == 0 {
				n = n / 2
			} else {
				n = 3*n + 1
			}
			c++
		}

	}

	return c, nil
}
