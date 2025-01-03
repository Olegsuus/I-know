package cycleDoubleLinkedList

import (
	"fmt"
	"strings"
)

type Node struct {
	Next     *Node
	Previous *Node
	Value    int
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

type DoubleCircularLinkedList struct {
	Head  *Node
	count int
}

func (list *DoubleCircularLinkedList) CetCount() int {
	return list.count
}

func (list *DoubleCircularLinkedList) Print(forward bool) string {
	if list.count == 0 {
		return ""
	}

	if forward {
		return list.print(list.Head, forward)
	}
	return list.print(list.Head.Previous, forward)
}

func (list *DoubleCircularLinkedList) print(node *Node, forward bool) string {
	if list.count == 0 {
		return ""
	}
	var result strings.Builder
	current := node

	for {
		result.WriteString(fmt.Sprintf("%d", current.Value))
		if forward {
			current = current.Next
		} else {
			current = current.Previous
		}
		if current == node {
			break
		}
	}
	return result.String()
}

func (list *DoubleCircularLinkedList) Find(key int) *Node {
	if list.count == 0 {
		return nil
	}
	current := list.Head

	for {
		if current.Value == key {
			return current
		}
		current = current.Next
		if current == list.Head {
			break
		}
	}
	return nil
}

func (list *DoubleCircularLinkedList) FindLast(key int) *Node {
	if list.count == 0 {
		return nil
	}
	var node *Node
	current := list.Head

	for {
		if current.Value == key {
			node = current
		}
		current = current.Next

		if current == list.Head {
			break
		}
	}
	return node
}

func (list *DoubleCircularLinkedList) addBeforeInternal(node, newNode *Node) {
	newNode.Next = node
	newNode.Previous = node.Previous
	node.Previous.Next = newNode
	node.Previous = newNode
	list.count++
}

func (list *DoubleCircularLinkedList) insertNodeToEmptyList(node *Node) {
	node.Next = node
	node.Previous = node
	list.Head = node
	list.count++
}

func (list *DoubleCircularLinkedList) AddBefore(node *Node, item int) {
	newNode := &Node{Value: item}

	list.addBeforeInternal(node, newNode)
	if node == list.Head {
		list.Head = newNode
	}
}

func (list *DoubleCircularLinkedList) PushBack(item int) {
	newNode := &Node{Value: item}
	if list.count == 0 {
		list.insertNodeToEmptyList(newNode)
	} else {
		list.addBeforeInternal(list.Head, newNode)
	}
}

func (list *DoubleCircularLinkedList) PushFront(item int) {
	newNode := &Node{Value: item}
	if list.count == 0 {
		list.insertNodeToEmptyList(newNode)
	} else {
		list.addBeforeInternal(list.Head, newNode)
		list.Head = newNode
	}
}

func (list *DoubleCircularLinkedList) AddAfter(node *Node, item int) {
	if node == nil || list.count == 0 {
		return
	}

	newNode := &Node{Value: item}
	newNode.Next = node.Next
	newNode.Previous = node
	node.Next.Previous = newNode
	node.Next = newNode

	list.count++
}

func (list *DoubleCircularLinkedList) RemoveNode(node *Node) {
	if list.count == 0 {
		return
	}
	if node.Next == node {
		list.Head = nil
	} else {
		node.Next.Previous = node.Previous
		node.Previous.Next = node.Next
		if list.Head == node {
			list.Head = node.Next
		}
	}
	list.count--
}

func (list *DoubleCircularLinkedList) RemoveFirst() {
	list.RemoveNode(list.Head)
}

func (list *DoubleCircularLinkedList) RemoveLast() {
	list.RemoveNode(list.Head.Previous)
}

func (list *DoubleCircularLinkedList) Remove(item int) bool {
	if list.count == 0 {
		return false
	}

	current := list.Head
	for {
		if current.Value == item {
			list.RemoveNode(current)
			return true
		}
		current = current.Next
		if current == list.Head {
			return false
		}
	}
}

func (list *DoubleCircularLinkedList) RemoveLastByValue(item int) bool {
	if list.count == 0 {
		return false
	}

	var node *Node
	current := list.Head
	for {
		if current.Value == item {
			node = current
		}
		current = current.Next
		if current == list.Head {
			break
		}
	}

	if node != nil {
		list.RemoveNode(node)
		return true
	}

	return false
}
