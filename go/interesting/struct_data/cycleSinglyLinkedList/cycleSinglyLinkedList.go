package cycleSinglyLinkedList

import (
	"fmt"
)

type Node struct {
	Next  *Node
	Value int
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

type SinglyCircularLinkedList struct {
	Tail  *Node
	count int
}

func (s *SinglyCircularLinkedList) GetCount() int {
	return s.count
}

func (s *SinglyCircularLinkedList) Print() string {
	if s.count == 0 {
		return ""
	}
	return s.PrintNode(s.Tail.Next)
}

func (s *SinglyCircularLinkedList) PrintNode(node *Node) string {
	if s.count == 0 {
		return ""
	}
	result := ""
	current := node

	for {
		result += fmt.Sprintf("%d", current.Value)
		current = current.Next

		if current == nil {
			break
		}
	}
	return result
}

func (s *SinglyCircularLinkedList) Find(key int) *Node {
	if s.count == 0 {
		return nil
	}

	current := s.Tail

	for {
		if current.Value == key {
			return current
		}
		current = current.Next

		if current == s.Tail {
			break
		}
	}

	return nil
}

func (s *SinglyCircularLinkedList) FindLast(item int) *Node {
	if s.Tail == nil {
		return nil
	}

	var curKey *Node
	head := s.Tail.Next
	current := head

	for {
		if current.Value == item {
			curKey = current
		}
		current = current.Next
		if current == head {
			break
		}
	}

	return curKey
}

func (s *SinglyCircularLinkedList) addAfterInternal(node, newNode *Node) {
	newNode.Next = node.Next
	node.Next = newNode
	s.count++
}

func (s *SinglyCircularLinkedList) insertNodeToEmptyList(node *Node) {
	node.Next = node
	s.Tail = node
	s.count++
}

func (s *SinglyCircularLinkedList) PushBack(item int) {
	newNode := &Node{Value: item}
	if s.count == 0 {
		s.insertNodeToEmptyList(newNode)
	} else {
		s.addAfterInternal(s.Tail, newNode)
		s.Tail = newNode
	}
}

func (s *SinglyCircularLinkedList) PushFront(item int) {
	newNode := &Node{Value: item}
	if s.count == 0 {
		s.insertNodeToEmptyList(newNode)
	} else {
		s.addAfterInternal(s.Tail, newNode)
	}
}

func (s *SinglyCircularLinkedList) AddAfter(node *Node, item int) {
	newNode := &Node{Value: item}
	s.addAfterInternal(node, newNode)
	if node == s.Tail {
		s.Tail = newNode
	}
}

func (s *SinglyCircularLinkedList) AddBefore(node *Node, item int) {
	newNode := &Node{Value: item}
	current := s.Tail

	for {
		if current.Next == node {
			current.Next = newNode
			newNode.Next = node
		}
		current = current.Next
		if current == s.Tail {
			break
		}
	}
	s.count++
}

func (s *SinglyCircularLinkedList) removeAfterNodeInternal(node *Node) {
	if s.count == 0 {
		return
	}
	if s.count == 1 {
		s.Tail = nil
	} else {
		if node.Next == s.Tail {
			s.Tail = node
		}
		node.Next = node.Next.Next
	}
	s.count--
}

func (s *SinglyCircularLinkedList) RemoveFirst() {
	s.removeAfterNodeInternal(s.Tail)
}

func (s *SinglyCircularLinkedList) RemoveNode(node *Node) {
	if s.count == 0 {
		return
	}
	current := node

	for current.Next != node {
		current = current.Next
	}
	s.removeAfterNodeInternal(current)
}

func (s *SinglyCircularLinkedList) RemoveLast() {
	s.RemoveNode(s.Tail)
}

func (s *SinglyCircularLinkedList) Remove(key int) bool {
	if s.count == 0 {
		return false
	}

	current := s.Tail.Next
	var node *Node

	for {
		if current.Value == key {
			node = current
			break
		}
		current = current.Next

		if current == s.Tail.Next {
			break
		}
	}

	if node != nil {
		s.RemoveNode(node)
		return true
	}
	return false
}

func (s *SinglyCircularLinkedList) RemoveLastKey(key int) bool {
	if s.count == 0 {
		return false
	}

	current := s.Tail.Next
	var node *Node

	for {
		if current.Value == key {
			node = current
		}

		current = current.Next

		if current == s.Tail.Next {
			break
		}
	}

	if node != nil {
		s.RemoveNode(node)
		return true
	}
	return false
}
