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
	
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if grid[i][j] != 'A' { continue }
			m := make(map[byte]int)
			inds := [][]int{{1,1},{1,-1},{-1,1},{-1,-1}}
			for _, ind := range inds {
				m[grid[i+ind[0]][j+ind[1]]]++
			}
			if m['M'] == 2 && m['S'] == 2 && grid[i+1][j+1] != grid[i-1][j-1] {
				ans++
			}
		}
	}
	
	fmt.Println(ans)
	
}