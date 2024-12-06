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
	output, _ := os.OpenFile("/home/sriteja/Competitive_Programming/Other Contests/Advent of Code/AOC2024-Go/output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	fmt.Fprintln(output, "Output")
	
	type Direction struct {
		x, y int
	}
	dirs := []Direction{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	
	var grid [][]byte
	var s string
	for _, err := fmt.Scan(&s); err == nil; _, err = fmt.Scan(&s) {
		grid = append(grid, []byte(s))
	}
	
	var x, y, dir int
	n, m := len(grid), len(grid[0])
	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			if grid[i][j] == '^' {
				x, y = i, j
				break
			}
		}
	}
	
	for x>=0 && x<n && y>=0 && y<m {
		if grid[x][y] == '#' {
			x -= dirs[dir].x
			y -= dirs[dir].y
			dir = (dir+1)%4
		}
		grid[x][y] = 'X'
		x += dirs[dir].x
		y += dirs[dir].y
	}
	ans := 0
	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			if grid[i][j] == 'X' {
				ans++
			}
		}
	}
	fmt.Println(ans)
	
}