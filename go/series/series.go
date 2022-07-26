package series

func All(n int, s string) []string {
	var list []string
	if n > len(s) {
		return nil
	} else if n == len(s) {
		list = append(list, s)
		return list
	}

	for i := 0; i <= len(s)-n; i++ {
		list = append(list, s[i:i+n])
	}

	return list
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	} else {
		return s[:n], true
	}
}
