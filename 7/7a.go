//go:build ignore

package main

import (
	"bufio"
	"fmt"
	// "io"
	// "math"
	"os"
	// "sort"
	"strconv"
	"strings"
	// "syscall"
)

func main() {
	output, _ := os.OpenFile("/home/sriteja/Competitive_Programming/Other Contests/Advent of Code/AOC2024-Go/output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	fmt.Fprintln(output, "Output")

	scanner := bufio.NewScanner(os.Stdin)
	var ans int64

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		a := make([]int64, len(fields)-1)
		for i, v := range fields[1:] {
			val, _ := strconv.Atoi(v)
			a[i] = int64(val)
		}
		tot, _ := strconv.Atoi(fields[0][:len(fields[0])-1])
		
		found := false
		
		var recur func(int, int64)
		recur = func(i int, sum int64) {
			if found { return }
			if i == len(a) {
				if sum == int64(tot) { found = true }
				return
			}
			recur(i+1, sum+a[i])
			recur(i+1, sum*a[i])
		}
		recur(1, a[0])
		if found { ans += int64(tot)}
	}
	fmt.Println(ans)
}
