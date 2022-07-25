package armstrong

import (
	"strconv"
)

func IsNumber(n int) bool {
	//convert to string
	sn := strconv.Itoa(n)
	an := 0
	l := len(sn)
	for _, v := range sn {
		an += pow(int(v-'0'), l)
	}

	return an == n
}

func pow(n1, n2 int) (v int) {
	v = 1
	for i := 0; i < n2; i++ {
		v *= n1
	}
	return
}
