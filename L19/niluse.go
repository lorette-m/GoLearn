package main

type StringStack struct {
	data []string // "zero" value ready-to-use
}

func (s *StringStack) Push(x string) {
	s.data = append(s.data, x)
}

func (s *StringStack) Pop() string {
	if l := len(s.data); l > 0 {
		t := s.data[l-1]
		s.data = s.data[:l-1]
		return t
	}

	panic("pop from empty stack")
}
