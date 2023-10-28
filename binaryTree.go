package binary_tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func BinaryTree() *Tree {
	return &Tree{}
}

func (tree *Tree) Insert(value int) {
	newNode := NewNode(value)
	if tree.Root == nil {
		tree.Root = newNode
		return
	}

	currentNode := tree.Root
	for {
		if value < currentNode.Value {
			if currentNode.Left == nil {
				currentNode.Left = newNode
				return
			}
			currentNode = currentNode.Left
		} else {
			if currentNode.Right == nil {
				currentNode.Right = newNode
				return
			}
			currentNode = currentNode.Right
		}
	}
}

func (t *Tree) Find(value int) bool {
	currentNode := t.Root
	for currentNode != nil {
		if value == currentNode.Value {
			return true
		} else if value < currentNode.Value {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	return false
}

func (t *Tree) Delete(index int) {
	currentNode, previousNode := t.Root, &Node{}
	for currentNode != nil {
		if index == currentNode.Value {
			if currentNode.Left != nil && currentNode.Right != nil {
				temp := currentNode.Left
				before := previousNode

				currentNode = currentNode.Right //RIGHT SUBTREE
				next := currentNode

				for currentNode.Left != nil {
					previousNode = currentNode
					if currentNode.Left != nil {
						currentNode = currentNode.Left
					} else {
						currentNode = currentNode.Right
					}
				}

				if currentNode.Right != nil {
					previousNode.Left = currentNode.Right
				} else {
					previousNode.Left = nil
				}

				currentNode.Right = next
				before.Left = currentNode
				currentNode.Left = temp
			} else if currentNode.Left == nil && currentNode.Right == nil {
				if previousNode.Left == currentNode {
					previousNode.Left = nil
				} else {
					previousNode.Right = nil
				}
			} else {
				if currentNode.Right != nil {
					previousNode.Left = currentNode.Right
				} else {
					previousNode.Left = currentNode.Left
				}
			}

			return
		}

		previousNode = currentNode
		if index < currentNode.Value {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
}

// func (t *Tree) FindNodes(value int) (*Node, *Node, error) {
// 	currentNode, previousNode := t.Root, &Node{}
// 	for currentNode != nil {
// 		if value == currentNode.Value {
// 			return currentNode, previousNode, nil
// 		}

// 		previousNode = currentNode
// 		if value < currentNode.Value {
// 			currentNode = currentNode.Left
// 		} else {
// 			currentNode = currentNode.Right
// 		}
// 	}
// 	return nil, nil, fmt.Errorf("Not exist")
// }

// func (t *Tree) Delete(index int) {
// 	currentNode, previousNode := t.Root, &Node{}

// 	for currentNode != nil {
// 		if index == currentNode.Value {

// 			temp := currentNode.Left         //49
// 			currentNode := currentNode.Right //55
// 			previousNode.Left = currentNode  //55

// 			for currentNode != nil {
// 				temp, currentNode.Left = currentNode.Left, temp //54 //49

// 				if temp != nil {
// 					temp.Right = currentNode
// 				}

// 				currentNode = currentNode.Right
// 			}

// 			return
// 		}

// 		previousNode = currentNode
// 		if index < currentNode.Value {
// 			currentNode = currentNode.Left
// 		} else {
// 			currentNode = currentNode.Right
// 		}
// 	}
// }
