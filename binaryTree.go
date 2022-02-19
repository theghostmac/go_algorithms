package main

import (
	"fmt"
)

// Tree type of recursive struct
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// traverse function
func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}
