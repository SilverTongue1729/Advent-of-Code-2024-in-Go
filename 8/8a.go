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
				x := 2 * posj.x - posi.x
				y := 2 * posj.y - posi.y
				if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) { continue }
				vis[x][y] = true
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