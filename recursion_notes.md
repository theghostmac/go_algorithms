# Recursion and Backtracking in Go
## Recursion
A function which calls itself is called recursive. It solves a problem by calling a 
copy of itself to work on smaller problems. It is important to make sure
that a recursion method terminates. 

Recursion is generally shorter and easier to write than iterative code. Even loops
get converted to recursive code by the compiler or interpreter. 
A recursive function performs a task in part by calling itself to perform 
the subtasks.

### Format of a Recursive Function
if (test for the base case) {
    return some base case value
} else if {test for another base case
    return some other base case value
} 
return (some work and then a recursive call)

