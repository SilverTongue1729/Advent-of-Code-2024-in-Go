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
	
	var sx, sy int
	n, m := len(grid), len(grid[0])
	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			if grid[i][j] == '^' {
				sx, sy = i, j
				grid[i][j] = '.'
				break
			}
		}
	}
	ans := 0
	
	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			if grid[i][j] != '.' { continue }
			grid[i][j] = '#'
			
			vis := make([][]int, n)
			for i:=0; i<n; i++ {
				vis[i] = make([]int, m)
			}
			
			flag := false
			x, y, dir := sx, sy, 0
			for x>=0 && x<n && y>=0 && y<m {
				if grid[x][y] == '#' {
					x -= dirs[dir].x
					y -= dirs[dir].y
					dir = (dir+1)%4
				}
				if (vis[x][y] & (1<<dir)) != 0 {
					flag = true
					break
				}
				vis[x][y] |= 1<<dir
				x += dirs[dir].x
				y += dirs[dir].y
			}
			
			if flag { ans++ }
			grid[i][j] = '.'
		}
	}
	
	fmt.Println(ans)
}