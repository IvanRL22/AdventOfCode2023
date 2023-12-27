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
## Day 6
### Part 1
The problem always will have a continuous range of solutions  
Therefore it can be solved by, for each time and distance, looking for the lowest and highest winning times  
After that, assume every time in between is also be a winner  
### Part 2
Read all numbers and join them to have the full time and distance  
Just remove iteration through all races, since there's just one  
Run the same logic just once  
Profit  
## Day 7
### [Part 1](Day7/Part1/day7_1.go)
This was supposed to be simple but turned ugly  
Had to "brute-force" it by finding rank, value by cards and then ordering  
Struggled with go orderings in general  
Realized after many tries that the full/double pair detection logic was flawed  
It was screwing the ordering on *some* comparisons, so the example data *almost* always turned the right result, but not the real data  
### [Part 2](Day7/Part2/day7_2.go)
Changed J to have the less value when calculating by card  
Added (messed up) logic for calculating rank with jokers if needed  
## Day 8
### [Part 1](Day8/Part1/day8_1.go)
Read node info and convert to map  
Iterate through instructions and navigate map accordingly until final node is found  