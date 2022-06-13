package table

import "fmt"

type Node struct {
	Key   int         // key = pointer_size - log2(list_size), like hashmap golang
	Value interface{} // Node's value
	Next  *Node       // Next node
}

type List struct {
	Head *Node
	Tail *Node
}

// Insert 1 node to tail of linked list
func (l *List) Insert(key int, value interface{}) {
	node := createNewNode(key, value)

	if l.checkEmpty() == true {
		// If empty linked list add to head
		l.Head = node
		l.Tail = node

	} else {

		// Not empty linked list
		l.Tail.Next = node
		l.Tail = node
	}

}

// Check empty linked list
func (l *List) checkEmpty() bool {
	// Return true if empty list (head = tail = nil)
	// Return false if empty
	if l.Head == nil && l.Tail == nil {
		return true
	} else {
		return false
	}
}

// Create 1 node with value k
func createNewNode(key int, value interface{}) *Node {
	node := &Node{
		Key:   key,
		Value: value,
		Next:  nil,
	}
	return node
}

// Display all node value in linked list
func (l *List) Display() {
	list := l.Head
	i := 0
	fmt.Println("-----Display list-------")

	for list != nil {
		fmt.Printf("%v : %v\n", i, list.Value)
		list = list.Next
		i++
	}
}

// Get list length
func (l *List) LengthOfList() int {
	node := l.Head
	count := 0

	for node != nil {
		count++
		node = node.Next
	}
	return count
}

// Get element at postion
// Return Node if found else return nil
func (l *List) GetElementAtPosition(position int) *Node {
	node := l.Head
	pos := 0

	for node != nil {
		if pos == position {
			return node
		}

		// Next position
		node = node.Next
		pos++
	}
	return nil
}
