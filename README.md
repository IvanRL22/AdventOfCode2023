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
### [Part 2](Day8/Part2/day8_2.go)
Couldn't have solved this without hints, never would have realized that LCM was the way to go  
## [Day 9](Day9/day9.go)
### Part 1
Used straight-forward approach with one minor hack
- Calculate each subsequent array but only store the last element (the only one we need)  
- Sum all the calculated last elements plus the last element from original data to get final row result  
### Part 2
Same approach, just tweaked algorithm to use first element of arrays and tweaked calculations  
## Day 10
## [Day 11](Day11/day11.go)
Started meddling with testing instead of reading from test file
### Part 1
I realized there was a simpler way to calculate than actually expanding input:  
- Save which rows/columns don't have galaxies  
- Distance is just the difference in column position plus the difference in row position (see [Manhattan distance](https://en.wikipedia.org/wiki/Taxicab_geometry))  
- Then check if there are empty rows or columns in between and add 1 for each  
### Part 2
This was trivial due to the way part 1 was calculated:  
- Just add 999.999 for every empty row/column  
- The reason for 999.999 instead of 1M is that we had already counted each row/column inbetween galaxies once  
## Day 12
### Part 1
Brute force for the win  