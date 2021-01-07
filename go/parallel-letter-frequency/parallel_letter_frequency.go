package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency concurrently counts the frequency of each rune in a given text and returns
// this data as a FreqMap.
func ConcurrentFrequency(xs []string) FreqMap {
	m := FreqMap{}
	for _, s := range xs {
		freq := make(chan FreqMap)
		go func(str string) {
			freq <- Frequency(str)
		}(s)
		for k, v := range <-freq {
			m[k] += v
		}
	}
	return m
}
