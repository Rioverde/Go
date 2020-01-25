/*

We aim to implement Depth First Search (DFS) on a grid, program will accept a problem text file and output a solution text file.

1. (0,6)(5,1)...(6,9) // List of available spaces in the grid.
2. 4                  // Number of problems in this file.
3. 10 11 9 0          // Problem #1, first pair is start position, second pair is end position.
4. 3 2 4 4            // Problem #2
5. 2 7 9 11           // Problem #3
6. 1 1 11 11          // Problem #4



Grid Rules:
From a position, you can only move
- UP    : 2 squares
- DOWN  : 2 squares
- LEFT  : 1 square
- Right : 3 squares


1. 10 (10,11)(9,11)(7,8)(6,6)...(9,0) // First integer stands for minimum number of moves 'n'.
2. 3  (3,2)...(4,4)                   // Following the first integer is the coordinates of
3. 9  (3,2)...(4,4)                   // the path taken to reach from start coordinate to
4. 15 (1,1)...(11,11)                 // end destiation

How to parse problem.txt

Step 1) Create a .go file take problem.txt as input.
Step 2) Read each character one-by-one using loops or other methods
Step 3) Filter the characters that are not needed.
Step 4) After filtering store them,so that they can be accessed later.
Step 5) Implement DFS.

*/

package main

import (
	grid "github.com/Rioverde/Go/Projects/AI/crud"
)

func main() {

	grid.GenerateGrid()

}
