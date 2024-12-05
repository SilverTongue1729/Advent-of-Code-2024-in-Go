//go:build ignore

package main

import (
	"fmt"
	"io"
	"sort"
	"math"
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
	
	for i := 0; i < n; i++ {
		ans += int(math.Abs(float64(a[i] - b[i])))
	}
	
	fmt.Println(ans)
}