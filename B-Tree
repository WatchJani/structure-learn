package b_tree

// here we can define deferent key structure for our database indexing system
// for example, our key can be multiple column index for name and age, [name, age]
// here is example of one key index
type Key struct {
	value int
	index []int //reference on data in database/file
}

func NewKey(value int) Key {
	return Key{
		value: value,
	}
}

type Node struct {
	capacity int
	key      []Key
	nextNode []*Node
	prevues  *Node
}

func (n *Node) UpdateCapacity() {
	n.capacity++
}

// insert sorted key
func InsertIndex(slice []Key, index int, insert Key) []Key {
	if len(slice)-1 < index {
		return append(slice, insert)
	}

	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = insert

	return slice
}

func (n *Node) Append(key, index int) {
	n.key = InsertIndex(n.key, index, NewKey(key))
}

func NewNode(degree int) *Node {
	return &Node{
		key:      make([]Key, 0, degree), //degree-1
		nextNode: make([]*Node, degree),
	}
}

type Tree struct {
	memory []*Node
	degree int
}

func NewBTree(degree, capacity int) *Tree {
	memory := make([]*Node, 0, capacity)     //optimization pre load memory
	memory = append(memory, NewNode(degree)) //create empty node

	return &Tree{
		degree: degree,
		memory: memory,
	}
}

func (t *Tree) Insert(key int) {
	root := t.memory[0]

	//search
	index := Search(root, key)

	root.AppendKey(key, index) // add new key

	if root.capacity < t.degree {
		//dovrsiti djeljenje
	}
}

// Search finds the position of the key in the B-tree node.
// It returns the position within the current node.
func Search(node *Node, key int) int {
	var index int

	for node != nil {
		index = FindKey(node.key, key)

		if node.nextNode[index] == nil {
			break
		}

		node = node.nextNode[index]
	}

	return index
}

// FindKey finds the next route for the next node. // need to be log(n) -> fix later
func FindKey(keys []Key, element int) int {
	var position = -1

	for index, key := range keys {
		if key.value >= element {
			break
		}
		position = index
	}

	return position + 1
}

func (t *Tree) AppendNode() {
	t.memory = append(t.memory, NewNode(t.degree))
}

func (n *Node) AppendKey(key, index int) {
	n.Append(key, index)
	n.UpdateCapacity()
}
