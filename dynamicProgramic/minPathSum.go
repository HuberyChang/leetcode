package main

import "fmt"

func minPathSum(grid [][]int) int {
	l := len(grid)
	if l < 1 {
		return 0
	}
	for i := 0; i < l; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j != 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if i != 0 && j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else if i != 0 && j != 0 {
				grid[i][j] = mingrid(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[l-1][len(grid[l-1])-1]
}

func mingrid(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func main() {
	grd := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grd)) 
}
