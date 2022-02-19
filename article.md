# Non-Linear Data Structures

## Trees
Trees are non-linear data structures used mainly for searching. Trees have maximum nodes of
two children and minimum of none. In trees, the property values of the left node is lesser than 
the property values of the right node. Nodes with no children are considered leaves.
### Binary search tree
Binary search trees allow look-up operations, as well as addition, and removal of elements.
They store keys in a sorted order, to make look-up operations faster. Their space complexity is 
of the order O(n), but their operations like insert, search, and delete are of the order O(log n).
Binary trees have the following properties:
- key of type int
- value of type int
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
### Implementing binary trees

