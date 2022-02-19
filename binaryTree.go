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

// traverse function allows visiting of nodes with recursion
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

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	}
	t.Right = insert(t.Right, v)
	return t
}