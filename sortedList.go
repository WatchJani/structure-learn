package sorted_list

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func New(value int, next *Node) *Node {
	return &Node{
		Value: value,
		Next:  next,
	}
}

func List() *Node {
	return &Node{}
}

func (n *Node) Inset(value int) {
	for node := n.Next; node != nil; node = node.Next {
		if node.Value <= value {
			node.Next = New(value, node.Next) //next node is the node who has the value from input, and get pointer on the next node in the row :D
			return
		}
	}

	//first node(just in this case)
	n.Next = &Node{
		Value: value,
	}
}

func (l *Node) Read() {
	for node := l.Next; node != nil; node = node.Next { // why l.Next node, because we have root node and we want to skip it
		fmt.Println(node.Value)
	}
}
func (l *Node) Delete(index int) {
	if index < 0 {
		return
	}

	if index == 0 {
		l.Next = l.Next.Next
		return
	}

	currentNode := l.Next // why l.Next node, because we have root node and we want to skip it
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
	}

	if currentNode.Next != nil {
		currentNode.Next = currentNode.Next.Next //if we check last node we can go forward because we have a root node, root node will never be deleted
	}
}
