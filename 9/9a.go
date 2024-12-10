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
	
	var s string
	fmt.Scan(&s)
	var a []int
	for _, ch := range s {
		a = append(a, int(ch-'0'))
	}
	
	ans := 0
	li, ri := 0, (len(s)-1) / 2 * 2
	lj, rj := 0, 0
	ind := 0
	for li <= ri {
		if li % 2 == 0 {
			cnt := 0
			if li == ri {
				cnt = rj
			}
			for i:=0; i<a[li] - cnt; i++ {
				ans += ind * (li / 2)
				ind++
			}
			li++
			continue
		}
		
		if lj == a[li] {
			li++
			lj = 0
			continue
		}
		
		if rj == a[ri] {
			ri-=2
			rj = 0
			continue
		}
		
		ans += ind * (ri / 2)
		ind++
		lj++
		rj++
	}
	
	fmt.Println(ans)
}