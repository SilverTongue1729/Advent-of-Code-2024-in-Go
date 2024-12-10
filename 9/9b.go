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

  sum := 0
	for i:=0; i<len(a); i++ {
		sum += a[i]
	}
	
	v := make([]int, sum)
	vind := make([]int, len(a))
	ind := 0
	
	for i:=0; i<len(a); i++ {
		vind[i] = ind
		for j:=0; j<a[i]; j++ {
			if i % 2 == 1 {
				v[ind] = -1
			} else {
				v[ind] = i/2
			}
			ind++
		}	
	}
			
	var st [10]int = [10]int{0}
	for i:=(len(a)-1) / 2 * 2; i>0; i-=2 {
		if a[i] == 0 { continue }
		cnt:=0
		for j:=st[a[i]]; j<vind[i]; j++ {
			if v[j] != -1 {
				cnt = 0
				continue;
			}
			cnt++
			if cnt == a[i] {
				for k:=0; k<a[i]; k++ {
					v[k+j-a[i]+1] = i/2
					v[k+vind[i]] = -1
				}
				st[a[i]] = j+1
				break
			}
		}
		if cnt != a[i] {
			st[a[i]] = vind[i]+a[i]
		}
		for j:=a[i]; j<=9; j++ {
			st[j] = max(st[j], st[a[i]])
		}
	}
	
	ans:=0
	for i:=0; i<len(v); i++ {
		if v[i] != -1 {
			ans += v[i] * i
		}
	}
	fmt.Println(ans)
}