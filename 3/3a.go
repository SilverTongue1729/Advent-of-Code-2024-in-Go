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
	
	var str, line string
	for _, err := fmt.Scan(&line); err == nil; _, err = fmt.Scan(&line) {
		// str += " " + line
		str += line
		fmt.Fprintf(os.Stderr, "line = %v\n", line)
	}
	fmt.Fprintf(os.Stderr, "str = %v\n", str)
	
	var ans int64
	n := len(str)
	for i := 0; i < n; i++ {
		if i+3 >= n { break }
		if str[i:i+4] != "mul(" { continue }
		i += 4
		if i >= n { break }
		
		flag := false
		var a, b int64
		for ; i < n; i++ {
			if str[i] == ',' { break }
			if str[i] < '0' || str[i] > '9' { flag = true; break }
			a = a*10 + int64(str[i]-'0')
		}
		if flag { continue }
		i++
		if i >= n { break }
		for ; i < n; i++ {
			if str[i] == ')' { break }
			if str[i] < '0' || str[i] > '9' { flag = true; break }
			b = b*10 + int64(str[i]-'0')
		}
		if flag { continue }
		ans += a*b
		fmt.Fprintln(os.Stderr, a, b, ans)
	}
	fmt.Println(ans)
}
