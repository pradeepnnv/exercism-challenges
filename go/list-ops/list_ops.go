package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	r := 0
	if s.Length() == 0 {
		return initial
	}
	for i, v := range s {
		if i != 0 {
			initial = r
		}
		r = fn(initial, v)

	}
	return r
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	r := 0
	if s.Length() == 0 {
		return initial
	}
	for i := s.Length() - 1; i >= 0; i-- {
		if i != s.Length()-1 {
			initial = r
		}
		r = fn(s[i], initial)
	}
	return r
}

func (s IntList) Filter(fn func(int) bool) IntList {
	nl := IntList(make([]int, 0, s.Length()))
	for i := range s {
		if fn(s[i]) {
			nl = nl.Append([]int{s[i]})
		}
	}
	return nl
}

func (s IntList) Length() int {
	l := 0
	for i := range s {
		i++
		l++
	}
	return l
}

func (s IntList) Map(fn func(int) int) IntList {
	nl := make([]int, s.Length())
	for i := range nl {
		nl[i] = fn(s[i])
	}
	return nl
}

func (s IntList) Reverse() IntList {
	rl := make([]int, s.Length())
	for i := 0; i < s.Length(); i++ {
		rl[s.Length()-i-1] = s[i]
	}
	return rl
}

func (s IntList) Append(lst IntList) IntList {
	rl := make([]int, s.Length()+lst.Length())
	for i := 0; i < s.Length(); i++ {
		rl[i] = s[i]
	}
	for i := 0; i < lst.Length(); i++ {
		rl[s.Length()+i] = lst[i]
	}
	return rl
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, l := range lists {
		s = s.Append(l)
	}
	return s
}
