//go:build ignore

package main

import (
	"fmt"
	"io"
	// "math"
	"sort"
)

func main() {
	var a, b []int
	for {
		var x, y int
		_, err := fmt.Scan(&x, &y)
		if err == io.EOF {
			break
		}
		a = append(a, x)
		b = append(b, y)
	}

	n := len(a)
	sort.Ints(a)
	sort.Ints(b)
	
	var ans int
	
	for i, j := 0, 0; i < n && j < n; {
		if a[i] == b[j] {
			ii := i
			for ; ii < n-1 && a[ii] == a[ii+1] ; ii++ { }
			jj := j
			for ; jj < n-1 && b[jj] == b[jj+1] ; jj++ { }
			ans += (ii - i + 1) * (jj - j + 1) * a[i]
			i = ii + 1
			j = jj + 1
			continue
		}
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		}
	}

	fmt.Println(ans)
}
