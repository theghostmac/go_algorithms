# Recursion and Backtracking in Go
## Recursion
A function which calls itself is called recursive. It solves a problem by calling a 
copy of itself to work on smaller problems. It is important to make sure
that a recursion method terminates. 

Recursion is generally shorter and easier to write than iterative code. Even loops
get converted to recursive code by the compiler or interpreter. 
A recursive function performs a task in part by calling itself to perform 
the subtasks.

## Notes on Recursion
1. Recursion has two cases: base case and recursive case
2. Recursion terminates at the base case
3. Iteration is more efficient than recursion due to overhead function calls
4. If iteration/recursion can solve it, then recursion/iteration can solve it too
5. Sometimes iterative algorithms do not exist for some problems. Recursion is suitable for these kinds of problems.

### Format of a Recursive Function
if (test for the base case) {
    return some base case value
} else if {test for another base case
    return some other base case value
} 
return (some work and then a recursive call)

## Recursion vs Iteration
1. Recursion terminates when a base case is reached, iteration terminates when a condition is proven as false.
2. Each recursive call requires extra space on the stack memory, each iteration does not require space.
3. If we get infinite recursion, the program runs out of memory and results in a stack overflow. If we get infinite loop, it will run forever as no memory is being created.
4. Some problems are easier solved with recursion, iteration are not clear code sometimes.

## Example of Recursive Algorithms
1. Fibonacci series
2. Factorial finding
3. Merge sort & Quick sort
4. Binary search
5. Tree Traversals and problems
6. Graph Traversals
7. Dynamic Programming
8. Divide and Conquer algorithms
9. Tower of Hanoi
10. Backtracking algorithms

# Recursion Problems and Solutions
## The Tower of Hanoi Problemg
Features: three towers, a number of disks with different diameters.
Algorithm: Let's name the towers *Source,* *Auxiliary,* and *Destination.*
1. Move the top (n-1)th disk from Source to Auxiliary tower.
2. Move the nth disk from Source to Destination tower.
3. Move the (n-1)th disk from 