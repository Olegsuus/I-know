package binary_tree

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) {
	if n == nil {
		return
	}
	if val < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: val}
		} else {
			n.Left.Insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

func (n *Node) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Printf("%d ", n.Value)
	n.Right.InOrder()
}

func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}
	if n.Value == val {
		return true
	}

	if val < n.Value {
		n.Left.Search(val)
	}

	return n.Right.Search(val)
}

func main() {
	n := &Node{Value: 10}

	nums := []int{1, 20, 30, 3, 5, 4}

	for _, v := range nums {
		n.Insert(v)
	}

	fmt.Printf("in orders travel: ")
	n.InOrder()
	fmt.Println()

	fmt.Println(n.Search(10))
}
