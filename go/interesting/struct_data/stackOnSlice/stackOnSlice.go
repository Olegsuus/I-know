package stackOnSlice

import (
	"errors"
	"fmt"
)

type StackOnSlice struct {
	items []int
	size  int
}

func NewStackOnSlice() *StackOnSlice {
	return &StackOnSlice{
		items: make([]int, 0),
		size:  0,
	}
}

func (s *StackOnSlice) Print() string {
	result := ""
	for i := s.size - 1; i >= 0; i-- {
		result += fmt.Sprintf("%d ", s.items[i])
	}
	return result
}

func (s *StackOnSlice) Empty() bool {
	return s.size == 0
}

func (s *StackOnSlice) increaseSlice() {
	newCount := s.size * 2
	if s.size == 0 {
		newCount = 4
	}

	newSlice := make([]int, s.size, newCount)
	copy(newSlice, s.items)
	s.items = newSlice
}

func (s *StackOnSlice) Push(item int) {
	if s.size == len(s.items) {
		s.increaseSlice()
	}
	s.items = append(s.items, item)
	s.size++
}

func (s *StackOnSlice) Pop() (int, error) {
	if s.size == 0 {
		return 0, errors.New("stack is empty")
	}
	s.size--
	return s.items[s.size], nil
}

func (s *StackOnSlice) Peek() (int, error) {
	if s.size == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.items[s.size-1], nil
}

func (s *StackOnSlice) Clear() {
	s.size = 0
}

func main() {
	stack := NewStackOnSlice()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Print())

	del, err := stack.Pop()
	fmt.Println(err)
	fmt.Println(del)

	stack.Clear()
	stack.Empty()
}
