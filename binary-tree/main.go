package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert
func (n *Node) Insert(key int) {
	if key < n.Key {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	} else if key > n.Key {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	}
}

// Search
func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}
	if key < n.Key {
		// move left
		return n.Left.Search(key)
	} else if key > n.Key {
		// move right
		return n.Right.Search(key)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(101)
	tree.Insert(102)
	fmt.Println(tree.Search(10))
}