package lsproduct

import (
	"fmt"
	"regexp"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span == 0 && digits == "" {
		return 1, nil
	} else if digits == "" || len(digits) < span || span < 0 {
		return 0, fmt.Errorf("invalid input")
	} else if span == 0 {
		return 1, nil
	}
	match, _ := regexp.MatchString(`[^0-9]`, digits)
	if match {
		return 0, fmt.Errorf("invalid input")
	}
	var product int64
	for i := 0; i <= len(digits)-span; i += 1 {
		// fmt.Println("slice worked on:", digits[i:i+span])
		pv := productValue(digits[i : i+span])
		// fmt.Println("product value:", pv)
		if product < pv {
			product = pv
		}
	}
	return product, nil
}

func productValue(s string) (v int64) {
	// fmt.Println(s)
	v = 1
	for _, c := range s {
		v = v * int64(c-'0')
		// fmt.Println(v)
	}
	return
}
