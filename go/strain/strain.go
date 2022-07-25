package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {

	var o Ints
	for _, v := range i {
		if filter(v) {
			o = append(o, v)
		}
	}
	return o
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var o Ints
	for _, v := range i {
		if !filter(v) {
			o = append(o, v)
		}
	}
	return o
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var o Lists
	for _, v := range l {
		if filter(v) {
			o = append(o, v)
		}
	}
	return o
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var o Strings
	for _, v := range s {
		if filter(v) {
			o = append(o, v)
		}
	}
	return o
}
