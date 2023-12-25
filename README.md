# AdventOfCode2023
A fun way to learn Go


## Day 5
### Part 1
Started by creating maps that had the calculations of al explicit ranges  
This was a valid solution for the test data, but computationlly too heavy for the input data  
Created some structs to hold the ranges data and be able to compute the mappings on the fly  
### Part 2
Applied range concept for seeds and brute forced it  
Would like to try and improve with a more efficient algowithm  
## Day6
### Part1
The problem always will have a continuous range of solutions  
Therefore it can be solved by, for each time and distance, looking for the lowest and highest winning times  
After that, assume every time in between is also be a winner  