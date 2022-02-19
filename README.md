# go_algorithms
Real World Data Structures and Algorithms in Go, as well as HackerRank challenges I solved
## Runtime Tips
1. Most Go lookup operations that use built-in types have constant time (O(1)). Hence, built-in types are faster than custom types and should be used where possible, unless you want full control.
2. Array operations are faster than map operations, but maps are more versatile.
3. For smaller datasets, it doesn't really matter the algorithm you choose.

## Binary tree properties
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

## Binary tree facts
1. Because binary trees represent data in hierarchy, they are suitable during compilation in programming languages.
2. Binary trees are already orderly. 
3. Deleting an element is hard. 
4. For very large datasets, the height is small, and traversing is fast.
5. 