//go:build ignore

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ans int
	
	for scanner.Scan() {
		line := scanner.Text()
		
		fields := strings.Fields(line)
		a := make([]int, len(fields))
		for i, v := range fields {
			a[i], _ = strconv.Atoi(v)
		}
		
		for i := 0; i < len(a)-1; i++ {
			if a[i] < a[i+1] && a[0] < a[1] && a[i]+1 <= a[i+1] && a[i]+3 >= a[i+1] {
				continue
			} 
			if a[i] > a[i+1] && a[0] > a[1] && a[i]-1 >= a[i+1] && a[i]-3 <= a[i+1] {
				continue
			}
			ans--;
			break;
		}
		ans++;
	}
	
	fmt.Println(ans)
}