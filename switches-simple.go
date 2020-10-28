//Algorithm no 4
package main

import (
	"fmt"
)

func main(){
	end := 100
	a := 1
	b := 3
	rn := 0

	for i := 0; i < end; i++ {
		rn += 1
		b = 3 + 2 * i
		a += b
		if a > end {
			break
		}
	}
	fmt.Println(rn)
	// 1, 4, 9, 16, 25, 36, 49, 64, 81, 100
}