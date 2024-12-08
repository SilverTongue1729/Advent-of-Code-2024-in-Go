//go:build ignore

package main

import (
	// "bufio"
	"fmt"
	// "io"
	// "math"
	"os"
	// "sort"
	// "strings"
	// "strconv"
	// "syscall"
)

func GCD(a, b int) int {
	if a < 0 { a = -a }
	if b < 0 { b = -b }

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	output, _ := os.OpenFile("/home/sriteja/Competitive_Programming/Other Contests/Advent of Code/AOC2024-Go/output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	fmt.Fprintln(output, "Output")
	
	var grid []string
	var s string
	for _, err := fmt.Scan(&s); err == nil; _, err = fmt.Scan(&s) {
		grid = append(grid, s)
	}
	
	type Position struct {
		x, y int
	}
	
	m := make(map[rune][]Position)
	
	for i, row := range grid {
		for j, ch := range row {
			if ch == '.' { continue }
			if _, ok := m[ch]; !ok {
				m[ch] = make([]Position, 0)
			}
			m[ch] = append(m[ch], Position{i, j})
		}
	}
	var vis [][]bool
	for i:=0; i<len(grid); i++ {
		vis = append(vis, make([]bool, len(grid[0])))
	}
	
	// iterate over each rune in m
	for _, positions := range m {
		for i, posi := range positions {
			for j, posj := range positions {
				if i == j { continue }
				x := posj.x - posi.x
				y := posj.y - posi.y
				g := GCD(x, y)
				x /= g
				y /= g
				
				for k:=1; ; k++ {
					nx, ny := posi.x + k*x, posi.y + k*y
					if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) { break }
					vis[nx][ny] = true
				}
			}
		}
	}
	
	ans := 0
	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++ {
			if vis[i][j] { ans++ }
		}
	}
	
	fmt.Println(ans)
}