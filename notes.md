# DATA STRUCTURES AND ALGORITHMS IN GO
## Algorithms
An algorithm is a step-by-step unambiguous instruction to solve 
a give problem.
### Analysis of Algorithms
Algorithmic analysis helps to determine which algorithm is efficient 
in terms of time and space consumed.
The goal of analysis is to determine the amount of time and space resources 
required to execute it. 
### Running time analysis
It is the process of determining how processing time increases as the
size of the problem input size increase.
## Comparing Algorithms
Comparison is not done in:
- execution time
- number of statements executed
But it is done by considering the functions and running times.
## Rate of Growth
The rate at which the running time increases as a function of input.
Just take the highest number after approximation.
    
n^4 + 2n^2 + 100n + 500  has the highest rate of growth is n^4.

## Commonly Used Rates of Growth
Time Complexity, name, description
1. O(1) - constant time - input size n
2. O(log n) - logarithmic time - slower growing than ever linear function
3. O(n) - linear time - functions grow linearly with the input size n
4. O(n log n) - linear logarithmic time - faster growing than linear but slower than quadratic
5. O(n^2) - quadratic time - they grow faster than the linear logarithm function
6. O(n^3) - cubic time - faster than quadratic but slower than exponential
7. O(2^n) - exponential - faster than all functions except factorial functions.
8. O(n!) - factorial - faster than all functions

## Types of Analysis
1. Worst case: defines the input for which the algorithm takes the longest time, input for which the algorithm runs the slowest.
2. Best case: defines the input for which the algorithm takes the least time, input for which the algorithm runs the fastest.
3. Average case: predicts the running time of the algorithm, runs the algorithm many times with different inputs, compute the total running time, and assumes the input is random.

Lower Bound <= Average Time <= Upper Bound

Worst case: f(n) = n^2 + 500
Best case: f(n) = n + 100n + 500

# Asymptotic Notation
They are used to represent upper and lower bounds.