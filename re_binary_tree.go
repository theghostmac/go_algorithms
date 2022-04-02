package main

import "sync"

type Node struct {
	key       int
	value     int
	leftNode  *Node
	rightNode *Node
}

type BinarySearchTree struct {
	lock     sync.RWMutex
	rootNode *Node
}

// binary trees have three operations
// - lookup
// - addition
// - removal

func (tree *BinarySearchTree) InsertElement(key, value int) {
	tree.lock.Lock()
}
