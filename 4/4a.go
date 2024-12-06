//go:build ignore

package main

import (
	"fmt"
	// "io"
	// "math"
	"os"
	// "sort"
	// "strings"
	// "syscall"
)

func main() {
	file, _ := os.OpenFile("/home/sriteja/Competitive_Programming/Other Contests/Advent of Code/AOC2024-Go/output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	os.Stderr = file
	
	var grid []string
	var s string
	for _, err := fmt.Scan(&s); err == nil; _, err = fmt.Scan(&s) {
		grid = append(grid, s)
	}
	fmt.Fprintln(os.Stderr, grid)
	
	var ans int
	n,m := len(grid), len(grid[0])
	
	// horizontal
	for i := 0; i < n; i++ {
		for j := 0; j < m-3; j++ {
			if grid[i][j:j+4] == "XMAS" { ans++ }
			if grid[i][j:j+4] == "SAMX" { ans++ }
		}
	}
	
	// vertical
	for i := 0; i < n-3; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 'X' && grid[i+1][j] == 'M' && grid[i+2][j] == 'A' && grid[i+3][j] == 'S' { ans++ }
			if grid[i][j] == 'S' && grid[i+1][j] == 'A' && grid[i+2][j] == 'M' && grid[i+3][j] == 'X' { ans++ }
		}
	}
	
	// diagonal top-left to bottom-right
	for i := 0; i < n-3; i++ {
		for j := 0; j < m-3; j++ {
			if grid[i][j] == 'X' && grid[i+1][j+1] == 'M' && grid[i+2][j+2] == 'A' && grid[i+3][j+3] == 'S' { ans++ }
			if grid[i][j] == 'S' && grid[i+1][j+1] == 'A' && grid[i+2][j+2] == 'M' && grid[i+3][j+3] == 'X' { ans++ }
		}
	}
	
	// diagonal top-right to bottom-left
	for i := 0; i < n-3; i++ {
		for j := 3; j < m; j++ {
			if grid[i][j] == 'X' && grid[i+1][j-1] == 'M' && grid[i+2][j-2] == 'A' && grid[i+3][j-3] == 'S' { ans++ }
			if grid[i][j] == 'S' && grid[i+1][j-1] == 'A' && grid[i+2][j-2] == 'M' && grid[i+3][j-3] == 'X' { ans++ }
		}
	}
	
	fmt.Println(ans)
	
}