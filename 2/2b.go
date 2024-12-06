//go:build ignore

package main

import (
	"bufio"
	"fmt"
	// "log"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	file, _ := os.OpenFile("/home/sriteja/Competitive_Programming/Other Contests/Advent of Code/AOC2024-Go/output.txt", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	_ = syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd()))

	scanner := bufio.NewScanner(os.Stdin)
	var ans int

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		a := make([]int, len(fields))
		for i, v := range fields {
			a[i], _ = strconv.Atoi(v)
		}
		n := len(a)
		
		verify := func(val1, val2 int, swap bool) bool {
			if swap { val1, val2 = val2, val1 }
			if val1 < val2 && val1+1 <= val2 && val2 <= val1+3 { return true }
			return false
		}
		
		for _, swap := range []bool{false, true} {
			flag := 2
			a1 := a
			for i := 0; i < n-1; i++ {
				if verify(a1[i], a1[i+1], swap) { continue }
				if flag == 2 {
					flag = 1
					if i == n-2 { continue }
					if verify(a1[i], a1[i+2], swap) {
						a1[i+1] = a1[i]
					} else if i > 0 && !verify(a1[i-1], a1[i+1], swap) { 
						flag = 0
						break
					}
					continue
				}
				flag = 0
				break
			}

			if flag == 0 { continue }
			ans++
			break
		}
		
	}

	fmt.Println(ans)
}
