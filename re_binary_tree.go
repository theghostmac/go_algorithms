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

// InsertElement inserts the element with a key and a value.
func (tree *BinarySearchTree) InsertElement(key, value int) {
	tree.lock.Lock()         // the tree's lock instance is locked first
	defer tree.lock.Unlock() // tree is unlocked before inserting an element
	var treeNode *Node
	treeNode = &Node{key, value, nil, nil}
	if tree.rootNode == nil { // if there is no existing root node,
		tree.rootNode = treeNode
	} else {
		insertTreeNode(tree.rootNode, treeNode) // implement the insertion of a new node
		// if there is an existing root Node
	}
}

// insertTreeNode is for addition and is tied to the root node
func insertTreeNode(rootNode *Node, newTreeNode *Node) {
	if newTreeNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.leftNode, newTreeNode)
		}
	} else {
		if rootNode.rightNode == nil {
			insertTreeNode(rootNode.leftNode, newTreeNode)
		} else {
			insertTreeNode(rootNode.rightNode, newTreeNode)
		}
	}
}

func (tree *BinarySearchTree) traverseInOrder(function func(int)) {
	tree.lock.RLock()         // RLock locks the tree from writing operations
	defer tree.lock.RUnlock() // RUnlock is deferred to open the tree after traversal

}