package blank_list

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func New(value int, node *Node) *Node {
	return &Node{
		Value: value,
		Next:  node,
	}
}

func List() *Node {
	return &Node{} //Root node
}

func (l *Node) Insert(value int) {
	l.Next = New(value, l.Next)
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