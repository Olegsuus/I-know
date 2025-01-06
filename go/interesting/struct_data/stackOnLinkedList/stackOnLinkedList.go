package stackOnLinkedList

import (
	"container/list"
	"errors"
	"fmt"
)

type StackOnLinkedList struct {
	items *list.List
}

func NewStackOnLinkedList() *StackOnLinkedList {
	return &StackOnLinkedList{
		items: list.New(),
	}
}

func (s *StackOnLinkedList) Print() string {
	var result string
	for elem := s.items.Back(); elem != nil; elem = elem.Prev() {
		result += fmt.Sprintf("%d ", elem.Value)
	}
	return result
}

func (s *StackOnLinkedList) Empty() bool {
	return s.items.Len() == 0
}

func (s *StackOnLinkedList) Push(item int) {
	s.items.PushBack(item)
}

func (s *StackOnLinkedList) Pop() (int, error) {
	if s.items.Len() == 0 {
		return 0, errors.New("stack is empty")
	}

	lastElement := s.items.Back()
	s.items.Remove(lastElement)
	return lastElement.Value.(int), nil
}

func (s *StackOnLinkedList) Peek() (int, error) {
	if s.items.Len() == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.items.Back().Value.(int), nil
}

func (s *StackOnLinkedList) Clear() {
	s.items.Init()
}

func main() {
	stack := NewStackOnLinkedList()

	fmt.Println(stack.Empty())

	stack.Push(1)
	stack.Push(2)

	fmt.Println(stack.Print())

	del, err := stack.Pop()
	fmt.Println(err == nil)
	fmt.Println(del)

	p, err := stack.Peek()
	fmt.Println(err == nil)
	fmt.Println(p)

	stack.Clear()
	fmt.Println(stack.Empty())
}
