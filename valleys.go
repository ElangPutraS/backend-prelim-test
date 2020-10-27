//Algorithm no 2
package main

import (
	"fmt"
)

func main(){
	var N int
    fmt.Scanf("%d", &N)
	acts := make([]string, N)
	height := 0
	res := 0
	isValley := false
    for i, _ := range acts {
		fmt.Scanf("%s", &acts[i])
		if acts[i] == "D" && height == 0 {
			isValley = true
		}
		if acts[i] == "U" && isValley && height == -1 {
			isValley = false
			res += 1
		} 
		if acts[i] == "D" {
			height -= 1
		} else {
			height += 1
		}
	}
	fmt.Println(res)
}