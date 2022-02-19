package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Tree type of recursive struct
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// traverse method allows visiting of nodes with recursion
func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

// create function populates the branches with random integers
func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}

// insert recursive function does lots of functions with each if statements
func insert(t *Tree, v int) *Tree {
	if t == nil { // if the tree is empty
		return &Tree{nil, v, nil}
	}
	if v == t.Value { // if the value exists in the tree
		return t
	}
	if v < t.Value { // if the value should go left or right
		t.Left = insert(t.Left, v)
	}
	t.Right = insert(t.Right, v)
	return t
}

// main function
func main() {
	tree := create(10)
	fmt.Println("The root of the tree is ", tree.Value, " \n.")
	traverse(tree)
	// fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	// fmt.Println()
	fmt.Println(" \n The root of the tree is ", tree.Value)
}
