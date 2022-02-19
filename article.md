# Data Structures and Algorithms with Go (2): Growing Your Go Skill with Trees

## Trees
Trees are non-linear data structures used mainly for searching. Trees have maximum nodes of
two children and minimum of none. In trees, the property values of the left node is lesser than 
the property values of the right node. Nodes with no children are considered leaves.

### Must-knows about trees
1. A binary tree is a data structure where each node has 0, 1,or 2 sub-nodes.
2. The root node is the first node of the tree, having no parents.
3. An edge is the link from child to parent.
4. The depth/height of the tree is the longest path from the root node to the furthest node.
5. The depth of a node is the number of edges from the node to the root of the tree.
6. A leaf is a node with no children.
7. The children of the same parent are called siblings.
8. Ancestors of a node are parent/grandparents to the parent of the node's parent.
9. Trees are either balanced or unbalanced, and balancing a tree is difficult and slow.
10. The size of a tree is the number of descendants it has including itself.

### Binary search tree
A tree is called a binary tree if it has a maximum of two child nodes. An empty tree 
is also a valid binary tree. Binary trees have a root node with a left and a right subtree.
Binary search trees allow look-up operations, as well as addition, and removal of elements.
They store keys in a sorted order, to make look-up operations faster. Their space complexity is 
of the order O(n), but their operations like insert, search, and delete are of the order O(log n).
Binary trees have the following properties:
- key of type ```int```
- value of type ```int```
- LeftNode of the node's type instance
- RightNode of the node's type instance

For example, the properties are arranged thus:
```go
type TreeNode struct {
	key int
	value int
	leftNode *TreeNode // best to put it first, as the orientation is left | value | right
	rightNode *TreeNode
}
```
### Binary tree operations
Operations possible with binary trees are grouped into two:

**Basic Operations**
- Inserting elements to a tree
- Deleting elements from a tree
- Searching for elements in a tree
- Traversing the tree

**Auxiliary Operations**
- Finding the size of a tree
- Finding the height of a tree
- Finding the level with the maximum sum
- Finding the least common ancestor (LCA) for a particular pair of nodes
- Others

### Applications of binary trees
The following are examples of use-cases of binary trees:
- Expression trees for compilers
- Huffman coding trees for data compression algorithms
- Binary search trees (BST) for searching, inserting, and deleting in O(log n)
- Priority queues (PQ) for searching and deleting minimum or maximum values in O(log n) time.

### Implementing binary trees in Go
Below is an example code where binary trees are implemented:
```go
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
```

## Conclusion
This article teaches the reader about Trees and binary trees. The possible operations 
with binary trees were also given, along with an implementation of some of these 
operations.
In a future article, applications of binary trees in problem-solving with the Go
programming languages will be given.