//go:build ignore

package main

import (
	"fmt"
	"strconv"
	"strings"
	// "io"
	// "math"
	"os"
	"sort"
	// "syscall"
)

func main() {
	file, _ := os.OpenFile("/home/sriteja/Competitive_Programming/Other Contests/Advent of Code/AOC2024-Go/output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	os.Stderr = file
	
	m := make(map[string]bool)
	var s string
	for _, err := fmt.Scanln(&s); err == nil; _, err = fmt.Scanln(&s) {
		if s == "" { break }
		m[s] = true
	}
	var ans int
	for _, err := fmt.Scan(&s); err == nil; _, err = fmt.Scan(&s) {
		a := strings.Split(s, ",")
		b := append([]string(nil), a...)
		
		sort.Slice(a, func(i, j int) bool {
			if m[a[i] + "|" + a[j]] {
				return true
			}
			return false
		})
		
		flag := true
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				flag = false
				break
			}
		}
		if flag { continue }
				
		num, _ := strconv.Atoi(a[len(a)/2])
		ans += num
	}
	fmt.Println(ans)
}