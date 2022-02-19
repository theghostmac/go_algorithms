# go_algorithms
Real World Data Structures and Algorithms in Go, as well as HackerRank challenges I solved
## Runtime Tips
1. Most Go lookup operations that use built-in types have constant time (O(1)). Hence, built-in types are faster than custom types and should be used where possible, unless you want full control.
2. Array operations are faster than map operations, but maps are more versatile.
3. For smaller datasets, it doesn't really matter the algorithm you choose.

## Binary Tree
1. A binary tree is a data structure where each node has 0, 1,or 2 sub-nodes.
2. The root node is the first node of the tree.
3. The depth/height of the tree is the longest path from the root node to the furthest node.
4. The depth of a node is the number of edges from the node to the root of the tree.
5. A leaf is a node with no children.
6. Trees are either balanced or unbalanced, and balancing a tree is difficult and slow.

