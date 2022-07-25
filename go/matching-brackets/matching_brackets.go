package brackets

import "fmt"

func Bracket(input string) bool {

	var s Stack
	for _, c := range input {
		switch c {
		case '{':
			s.Push(c)
		case '[':
			s.Push(c)
		case '(':
			s.Push(c)
		case ')':
			v, err := s.Pop()
			if err != nil || v != '(' {
				return false
			}
		case ']':
			v, err := s.Pop()
			if err != nil || v != '[' {
				return false
			}
		case '}':
			v, err := s.Pop()
			if err != nil || v != '{' {
				return false
			}
		default:
		}
	}

	return s.len == 0
}

type Stack struct {
	stack []rune
	len   int
}

func (s *Stack) Pop() (rune, error) {
	if s.len == 0 {
		return 0, fmt.Errorf("stack is empty...nothing to pop")
	}
	v := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	s.len--
	return v, nil
}

func (s *Stack) Push(r rune) {
	s.stack = append(s.stack, r)
	s.len++
}
