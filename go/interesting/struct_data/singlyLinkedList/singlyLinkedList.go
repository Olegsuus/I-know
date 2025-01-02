package singlyLinkedList

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

type SinglyLinkedList struct {
	Head *Node

	count int
}

func NewSinglyLinkedList(head *Node) *SinglyLinkedList {
	return &SinglyLinkedList{
		Head: head,
	}
}

func (l *SinglyLinkedList) GetCount() int {
	return l.count
}

func (l *SinglyLinkedList) Print() string {
	var result string
	current := l.Head

	for current != nil {
		result += fmt.Sprintf("%d ", current.Value)
		current = current.Next
	}

	return result
}

func (l *SinglyLinkedList) Find(key int) *Node {
	if l == nil || l.count == 0 {
		return nil
	}

	current := l.Head
	for current.Value != key {
		current = current.Next
		if current == nil {
			return nil
		}
	}
	return current
}

func (l *SinglyLinkedList) FindByIndex(index int) *Node {
	if l == nil || l.count == 0 {
		return nil
	}

	var i int
	current := l.Head
	for {
		if i == index {
			return current
		}
		if current.Next == nil {
			return nil
		}
		current = current.Next
		i++
	}
}

func (l *SinglyLinkedList) FindLast(key int) *Node {
	var curKey *Node

	current := l.Head
	for current != nil {
		if current.Value == key {
			curKey = current
		}
		current = current.Next
	}
	return curKey
}

func (l *SinglyLinkedList) PushBack(item int) {
	newNode := &Node{Value: item}

	if l.count == 0 {
		l.Head = newNode
	} else {
		tail := l.FindTail()
		tail.Next = newNode
	}
	l.count++
}

func (l *SinglyLinkedList) FindTail() *Node {
	if l.count == 0 {
		return nil
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	return current
}

func (l *SinglyLinkedList) PushFront(item int) {
	newNode := &Node{Value: item}

	if l.count != 0 {
		newNode.Next = l.Head
	}
	l.Head = newNode
	l.count++
}

func (l *SinglyLinkedList) AddAfter(node *Node, item int) {
	if node == nil {
		return
	}
	newNode := &Node{Value: item}
	newNode.Next = node.Next
	node.Next = newNode
	l.count++
}

func (l *SinglyLinkedList) AddBefore(node *Node, item int) {
	if node == nil {
		return
	}

	newNode := &Node{Value: item}
	current := l.Head

	if current == node {
		newNode.Next = current
		l.Head = newNode
		l.count++
		return
	}

	for current.Next != nil {
		if current.Next == node {
			current.Next = newNode
			newNode.Next = node
		}
		current = current.Next
	}
	l.count++
}

func (l *SinglyLinkedList) PushBackRange(items []int) {
	for i := 0; i < len(items); i++ {
		l.PushBack(items[i])
	}
}

func (l *SinglyLinkedList) RemoveFirst() {
	if l.count == 0 {
		return
	}

	l.Head = l.Head.Next
	l.count--
}

func (l *SinglyLinkedList) RemoveLast() {
	if l.count == 0 {
		return
	}

	if l.count == 1 {
		l.Head = nil
	} else {
		current := l.Head
		for current.Next.Next != nil {
			current = current.Next
		}
		current.Next = nil
	}
	l.count--
}

func (l *SinglyLinkedList) RemoveNode(node *Node) {
	if l.Head == node {
		l.Head = node.Next
	} else {
		current := l.Head
		for current.Next != nil {
			if current.Next == node {
				break
			}
			current = current.Next
		}
		current.Next = node.Next
	}
	l.count--
}

func (l *SinglyLinkedList) Remove(key int) bool {
	if l.count == 0 {
		return false
	}

	current := l.Head
	for current != nil {
		if current.Value == key {
			l.RemoveNode(current)
			return true
		}
		current = current.Next
	}
	return false
}

func (l *SinglyLinkedList) RemoveLastKey(key int) bool {
	if l.count == 0 {
		return false
	}

	lastKey := l.FindLast(key)
	if lastKey == nil {
		return false
	}

	l.RemoveNode(lastKey)

	return true
}

func (l *SinglyLinkedList) RemoveAll(key int) int {
	if l.count == 0 {
		return 0
	}

	isDelete := true
	var count int
	current := l.Head

	for {
		if isDelete = l.Remove(key); isDelete == true {
			count++
		} else {
			break
		}
		current = current.Next
	}
	return count
}

func (l *SinglyLinkedList) Reverse() {
	if l.count < 2 {
		return
	}

	var prev, next *Node
	current := l.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}

func main() {
	s := NewSinglyLinkedList(nil)
	s.PushBack(6)
	s.PushBack(7)
	s.PushFront(4)
	s.PushFront(3)

	s.PushBackRange([]int{8, 8, 8})

	fmt.Println(s.Print())
	s.Reverse()
	fmt.Println(s.Print())
}
