package stack

import "fmt"

type Stack[T rune] struct {
	data []T
}

func New[T rune]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Pop() T {
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Display() {
	for i := len(s.data) - 1; i >= 0; i-- {
		fmt.Println(string(s.data[i]))
	}
}
