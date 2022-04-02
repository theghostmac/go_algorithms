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
